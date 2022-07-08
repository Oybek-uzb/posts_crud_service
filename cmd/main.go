package main

import (
	"fmt"
	"github.com/Oybek-uzb/posts_crud_service/config"
	"github.com/Oybek-uzb/posts_crud_service/internal/services"
	pbp "github.com/Oybek-uzb/posts_crud_service/pkg/api/posts_crud_service"
	ser "github.com/Oybek-uzb/posts_crud_service/services"
	"google.golang.org/grpc"
	"net"
)

func main() {
	cfg := config.Load()

	grpcClients, _ := ser.NewGrpcClients(&cfg)

	s := grpc.NewServer()
	postsCRUDService := services.NewPostsCRUDService(grpcClients)
	pbp.RegisterPostsCRUDServiceServer(s, postsCRUDService)

	listen, err := net.Listen("tcp", cfg.HttpPort)
	if err != nil {
		return
	}
	fmt.Printf("Listening HTTP on %s\n", cfg.HttpPort)

	err = s.Serve(listen)
	if err != nil {
		return
	}
}
