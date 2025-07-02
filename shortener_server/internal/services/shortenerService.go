package services

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/base64"
	"fmt"
	"regexp"
	"time"

	"github.com/redis/go-redis/v9"
)

type ShortenerService struct {
	db    *sql.DB
	redis *redis.Client
}

func NewShortenerService(db *sql.DB, redis *redis.Client) *ShortenerService {
	return &ShortenerService{
		db:    db,
		redis: redis,
	}
}

func (s *ShortenerService) CreateShortUrl(url, userId string, startIndex int) (string, error) {
	if startIndex > 10 {
		// max attempt exceed for 10 times
		return "", fmt.Errorf("cannot generate short code. exceeded max 10 attempts")
	}

	byteURLData := []byte(url + fmt.Sprintf("::user_id=%s", userId))
	hashedURLData := fmt.Sprintf("%x", md5.Sum(byteURLData))
	shortCodeRegex, err := regexp.Compile("[/+]")
	if err != nil {
		return "", fmt.Errorf("unable to generate short code")
	}

	shortCodeData := shortCodeRegex.ReplaceAllString(base64.URLEncoding.EncodeToString([]byte(hashedURLData)), "_")
	if len(shortCodeData) < (startIndex + 6) {
		return "", fmt.Errorf("unable to generate short code")
	}

	shortCode := shortCodeData[startIndex : startIndex+6]

	// checking if the short code already exists
	query := "SELECT 1 FROM urls WHERE short = $1"
	var existing int
	if err := s.db.QueryRow(query, shortCode).Scan(&existing); err != nil && err != sql.ErrNoRows {
		return "", err
	}

	if existing != 0 {
		// already existing short code, recursion for further set
		return s.CreateShortUrl(url, userId, startIndex+1)
	}

	// setting to redis
	if err := s.redis.Set(context.Background(), shortCode, url, 0).Err(); err != nil {
		return "", fmt.Errorf("cannot set to redis: %s", err.Error())
	}

	// inserting to db
	insertQuery := "INSERT INTO urls(user_id, url, short) VALUES($1, $2, $3)"
	if _, err := s.db.Exec(insertQuery, userId, url, shortCode); err != nil {
		s.redis.Del(context.Background(), shortCode)
		return "", fmt.Errorf("cannot insert in db: %s", err.Error())
	}

	return shortCode, nil
}

func (s *ShortenerService) DeleteUrl(shortCode string) error {
	delUrlQuery := "DELETE FROM urls WHERE short = $1 RETURNING id"
	var deletedId int
	if err := s.db.QueryRow(delUrlQuery, shortCode).Scan(&deletedId); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no url found for this user")
		}
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.redis.Del(ctx, shortCode)

	if deletedId != 0 {
		delHistoryQuery := "DELETE FROM history WHERE url_id = $1"
		s.db.Exec(delHistoryQuery, deletedId)
	}

	return nil
}

func (s *ShortenerService) RegisterClick(shortCode string) error {
	query := "UPDATE urls SET clicks = clicks + 1 WHERE short = $1 RETURNING id"
	var urlId int
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	if err := tx.QueryRow(query, shortCode).Scan(&urlId); err != nil {
		tx.Rollback()
		return err
	}

	insertHistoryQuery := "INSERT INTO history(url_id) VALUES($1)"
	if _, err := tx.Exec(insertHistoryQuery, urlId); err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}
