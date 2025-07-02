package main

import (
	"log"
	"net"

	"github.com/Aniketyadav44/s7r/getter/internal/config"
	"github.com/Aniketyadav44/s7r/getter/internal/handlers"
	"github.com/Aniketyadav44/s7r/getter/internal/services"
	pb "github.com/Aniketyadav44/s7r/getter/proto"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("error in loading config: ", err.Error())
	}

	getterService := services.NewGetterService(cfg.Db)
	getterServer := handlers.NewGetterServer(getterService)

	net, err := net.Listen("tcp", ":5002")
	if err != nil {
		log.Fatal("error starting in a network: ", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGetterServer(grpcServer, getterServer)

	log.Println("starting getter service...")
	if err := grpcServer.Serve(net); err != nil {
		log.Fatal("error starting grpc server: ", err.Error())
	}
}
