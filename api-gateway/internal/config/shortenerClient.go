package config

import (
	"fmt"

	pb "github.com/Aniketyadav44/s7r/api-gateway/proto/shortener"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func loadShortenerGrpcClient(url string) (*pb.ShortenerClient, error) {
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("error in connecting to shortener grpc: %s", err.Error())
	}
	client := pb.NewShortenerClient(conn)
	return &client, nil
}
