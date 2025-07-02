package main

import (
	"log"
	"net"

	"github.com/Aniketyadav44/s7r/shortener_server/internal/config"
	"github.com/Aniketyadav44/s7r/shortener_server/internal/handlers"
	"github.com/Aniketyadav44/s7r/shortener_server/internal/services"
	pb "github.com/Aniketyadav44/s7r/shortener_server/proto"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("error in loading config: ", err.Error())
	}

	shortenerService := services.NewShortenerService(cfg.Db, cfg.Redis)
	shortenerServer := handlers.NewShortenerServer(shortenerService)

	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatal("error in starting tcp server: ", err.Error())
	}
	grpcServer := grpc.NewServer()
	pb.RegisterShortenerServer(grpcServer, shortenerServer)

	log.Println("starting shortener service...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("error in starting a new grpc server: ", err.Error())
	}
}
