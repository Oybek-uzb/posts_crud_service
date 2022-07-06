package services

import (
	"context"
	pcs "github.com/Oybek-uzb/posts_crud_service/pkg/api/posts_crud_service"
)

type PostsCRUDServiceServer struct {
}

func (s *PostsCRUDServiceServer) GetAllPosts(ctx context.Context, req *pcs.GetAllPostsRequest) (*pcs.GetAllPostsResponse, error) {

}
func (s *PostsCRUDServiceServer) GetPost(ctx context.Context, req *pcs.GetPostRequest) (*pcs.GetPostResponse, error) {

}
func (s *PostsCRUDServiceServer) UpdatePartialPost(ctx context.Context, req *pcs.UpdatePartialPostRequest) (*pcs.UpdatePartialPostResponse, error) {

}
func (s *PostsCRUDServiceServer) DeletePost(ctx context.Context, req *pcs.DeletePostRequest) (*pcs.DeletePostResponse, error) {
	
}
