package config

import (
	gpb "github.com/Aniketyadav44/s7r/api-gateway/proto/getter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func loadGetterGrpcClient(getterUrl string) (*gpb.GetterClient, error) {
	conn, err := grpc.NewClient(getterUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := gpb.NewGetterClient(conn)
	return &client, nil
}
