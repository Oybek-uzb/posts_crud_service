package services

import (
	"fmt"

	"github.com/Oybek-uzb/posts_crud_service/config"
	"github.com/Oybek-uzb/posts_crud_service/pkg/api/posts_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManager interface {
	PostsService() posts_service.PostsServiceClient
}

type grpcClients struct {
	postsService posts_service.PostsServiceClient
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	postsService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PostsServiceHost, conf.PostsServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		postsService: posts_service.NewPostsServiceClient(postsService),
	}, nil
}

func (g grpcClients) PostsService() posts_service.PostsServiceClient {
	return g.postsService
}
