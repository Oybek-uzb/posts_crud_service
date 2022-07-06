package main

import (
	"github.com/Oybek-uzb/posts_crud_service/config"
	"github.com/Oybek-uzb/posts_crud_service/internal/services"
	pbp "github.com/Oybek-uzb/posts_crud_service/pkg/api/posts_crud_service"
	"google.golang.org/grpc"
	"net"
)

func main() {
	cfg := config.Load()

	s := grpc.NewServer()
	postsCRUDService := services.NewPostsCRUDService()
	pbp.RegisterPostsCRUDServiceServer(s, postsCRUDService)

	listen, err := net.Listen("tcp", cfg.HttpPort)
	if err != nil {
		return
	}

	s.Serve(listen)

}
