package handlers

import (
	"context"

	"github.com/Aniketyadav44/s7r/shortener_server/internal/services"
	pb "github.com/Aniketyadav44/s7r/shortener_server/proto"
)

type ShortenerServer struct {
	pb.UnimplementedShortenerServer
	service *services.ShortenerService
}

func NewShortenerServer(service *services.ShortenerService) *ShortenerServer {
	return &ShortenerServer{
		service: service,
	}
}

func (h *ShortenerServer) ShortenUrl(c context.Context, req *pb.ShortenRequest) (*pb.ShortenResponse, error) {
	shortCode, err := h.service.CreateShortUrl(req.Url, req.UserId, 0)
	if err != nil {
		return &pb.ShortenResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.ShortenResponse{
		Success:   true,
		ShortCode: shortCode,
		Message:   "Short code created successfully",
	}, nil
}

func (h *ShortenerServer) DeleteUrl(c context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := h.service.DeleteUrl(req.ShortCode)
	if err != nil {
		return &pb.DeleteResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}
	return &pb.DeleteResponse{
		Success: true,
		Message: "Delete successfully",
	}, nil
}

func (h *ShortenerServer) RegisterClick(c context.Context, req *pb.RegisterClickRequest) (*pb.RegisterClickResponse, error) {
	err := h.service.RegisterClick(req.ShortCode)
	if err != nil {
		return &pb.RegisterClickResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}
	return &pb.RegisterClickResponse{
		Success: true,
		Message: "Click registered successfully",
	}, nil
}
