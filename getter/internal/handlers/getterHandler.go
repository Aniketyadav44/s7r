package handlers

import (
	"context"

	"github.com/Aniketyadav44/s7r/getter/internal/services"
	pb "github.com/Aniketyadav44/s7r/getter/proto"
)

type GetterServer struct {
	pb.UnimplementedGetterServer
	service *services.GetterService
}

func NewGetterServer(service *services.GetterService) *GetterServer {
	return &GetterServer{
		service: service,
	}
}

func (h *GetterServer) GetAllUrls(c context.Context, req *pb.AllRequest) (*pb.AllResponse, error) {
	urlItems, err := h.service.GetAllUrls(req.UserId)
	if err != nil {
		return &pb.AllResponse{
			Success: false,
			Message: err.Error(),
			Urls:    nil,
		}, nil
	}

	return &pb.AllResponse{
		Success: true,
		Message: "Success",
		Urls:    urlItems,
	}, nil
}

func (h *GetterServer) GetClicks(c context.Context, req *pb.ClickRequest) (*pb.ClickResponse, error) {
	count, err := h.service.GetClicks(req.ShortCode)
	if err != nil {
		return &pb.ClickResponse{
			Success: false,
			Message: err.Error(),
			Count:   0,
		}, nil
	}

	return &pb.ClickResponse{
		Success: true,
		Message: "Success",
		Count:   count,
	}, nil
}

func (h *GetterServer) GetHistory(c context.Context, req *pb.HistoryRequest) (*pb.HistoryResponse, error) {
	history, err := h.service.GetHistory(req.ShortCode)
	if err != nil {
		return &pb.HistoryResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.HistoryResponse{
		Success: true,
		Message: "Success",
		History: history,
	}, nil
}
