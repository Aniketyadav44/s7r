package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	pb "github.com/Aniketyadav44/s7r/api-gateway/proto/shortener"
	"github.com/redis/go-redis/v9"
)

type ShortenerService struct {
	redis          *redis.Client
	shortnerClient pb.ShortenerClient
}

func NewShortenerService(redis *redis.Client, shortenerClient *pb.ShortenerClient) *ShortenerService {
	return &ShortenerService{
		redis:          redis,
		shortnerClient: *shortenerClient,
	}
}

// picks url from short code via redis
func (s *ShortenerService) GetUrlFromShort(short string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	url, err := s.redis.Get(ctx, short).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return "", fmt.Errorf("invalid short code: no url found for code %s", short)
		}
		return "", err
	}

	return url, nil
}

// registers an url click via srhotener grpc service
func (s *ShortenerService) RegisterUrlClick(short string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := s.shortnerClient.RegisterClick(ctx, &pb.RegisterClickRequest{ShortCode: short})
	if err != nil || !res.Success {
		eMsg := err.Error()
		if res != nil {
			eMsg = res.Message
		}
		return errors.New(eMsg)
	}
	return nil
}

// creates a new short code via shortener grpc service
func (s *ShortenerService) CreateShortCode(url, userId string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := s.shortnerClient.ShortenUrl(ctx, &pb.ShortenRequest{Url: url, UserId: userId})
	if err != nil || !res.Success {
		eMsg := err.Error()
		if res != nil {
			eMsg = res.Message
		}
		return "", errors.New(eMsg)
	}

	return res.ShortCode, nil
}

// deletes a short code via shortener grpc service
func (s *ShortenerService) DeleteUrl(shortCode string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.shortnerClient.DeleteUrl(ctx, &pb.DeleteRequest{ShortCode: shortCode})
	if err != nil || !res.Success {
		eMsg := err.Error()
		if res != nil {
			eMsg = res.Message
		}
		return errors.New(eMsg)
	}
	return nil
}
