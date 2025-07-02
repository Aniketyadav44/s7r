package services

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/Aniketyadav44/s7r/api-gateway/internal/models"
	pb "github.com/Aniketyadav44/s7r/api-gateway/proto/getter"
)

type GetterService struct {
	getterClient pb.GetterClient
}

func NewGetterService(client *pb.GetterClient) *GetterService {
	return &GetterService{
		getterClient: *client,
	}
}

func (s *GetterService) GetAllUrls(userId string) ([]*models.URL, error) {
	uId, err := strconv.Atoi(userId)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := s.getterClient.GetAllUrls(ctx, &pb.AllRequest{UserId: int64(uId)})
	if err != nil || !res.Success {
		eMsg := err.Error()
		if res != nil {
			eMsg = res.Message
		}
		return nil, errors.New(eMsg)
	}

	urls := make([]*models.URL, 0)
	for _, v := range res.Urls {
		urls = append(urls, &models.URL{
			Id:        v.Id,
			UserId:    strconv.Itoa(int(v.UserId)),
			Url:       v.Url,
			Short:     v.Short,
			Clicks:    int(v.Clicks),
			CreatedAt: v.CreatedAt.AsTime(),
			UpdatedAt: v.UpdatedAt.AsTime(),
		})
	}
	return urls, nil
}

func (s *GetterService) GetClicks(shortCode string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := s.getterClient.GetClicks(ctx, &pb.ClickRequest{ShortCode: shortCode})
	if err != nil || !res.Success {
		eMsg := err.Error()
		if res != nil {
			eMsg = res.Message
		}
		return 0, errors.New(eMsg)
	}
	return int(res.Count), nil
}

func (s *GetterService) GetHistory(shortCode string) ([]time.Time, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := s.getterClient.GetHistory(ctx, &pb.HistoryRequest{ShortCode: shortCode})
	if err != nil || !res.Success {
		eMsg := err.Error()
		if res != nil {
			eMsg = res.Message
		}
		return nil, errors.New(eMsg)
	}

	timestamps := make([]time.Time, 0)
	for _, v := range res.History {
		timestamps = append(timestamps, v.AsTime())
	}
	return timestamps, nil
}
