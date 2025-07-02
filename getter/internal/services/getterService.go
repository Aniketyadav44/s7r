package services

import (
	"database/sql"
	"time"

	pb "github.com/Aniketyadav44/s7r/getter/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetterService struct {
	db *sql.DB
}

func NewGetterService(db *sql.DB) *GetterService {
	return &GetterService{
		db: db,
	}
}

func (s *GetterService) GetAllUrls(userId int64) ([]*pb.UrlItem, error) {
	query := "SELECT id, user_id, url, short, clicks, created_at, updated_at FROM urls where user_id = $1"

	rows, err := s.db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	urlItems := make([]*pb.UrlItem, 0)
	for rows.Next() {
		var url pb.UrlItem
		var createdAt time.Time
		var updatedAt time.Time
		if err := rows.Scan(&url.Id, &url.UserId, &url.Url, &url.Short, &url.Clicks, &createdAt, &updatedAt); err != nil {
			return nil, err
		}
		url.CreatedAt = timestamppb.New(createdAt)
		url.UpdatedAt = timestamppb.New(updatedAt)
		urlItems = append(urlItems, &url)
	}
	return urlItems, nil
}

func (s *GetterService) GetClicks(shortCode string) (int64, error) {
	query := "SELECT clicks FROM urls WHERE short = $1"
	var count int64
	if err := s.db.QueryRow(query, shortCode).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (s *GetterService) GetHistory(shortCode string) ([]*timestamppb.Timestamp, error) {
	query := "SELECT h.clicked_at FROM urls u LEFT JOIN history h ON h.url_id = u.id WHERE u.short = $1 ORDER BY h.clicked_at"
	rows, err := s.db.Query(query, shortCode)
	if err != nil {
		return nil, err
	}

	d := make([]*timestamppb.Timestamp, 0)
	for rows.Next() {
		var clickedAt time.Time
		if err := rows.Scan(&clickedAt); err != nil {
			return nil, err
		}
		d = append(d, timestamppb.New(clickedAt))
	}
	return d, nil
}
