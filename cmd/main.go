package main

import (
	"github.com/Oybek-uzb/posts_crud_service/internal/services"
	pbp "github.com/Oybek-uzb/posts_crud_service/pkg/api/posts_crud_service"
	"google.golang.org/grpc"
	"net"
)

func main() {
	s := grpc.NewServer()
	postsCRUDService := services.NewPostsCRUDService()
	pbp.RegisterPostsCRUDServiceServer(s, postsCRUDService)

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		return
	}

	s.Serve(listen)

}
