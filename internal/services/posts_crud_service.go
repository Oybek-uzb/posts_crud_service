package services

import (
	"context"
	"fmt"
	pbp "github.com/Oybek-uzb/posts_crud_service/pkg/api/posts_crud_service"
	pcs "github.com/Oybek-uzb/posts_crud_service/pkg/api/posts_crud_service"
)

type postsCRUDService struct {
	pbp.UnimplementedPostsCRUDServiceServer
}

func NewPostsCRUDService() *postsCRUDService {
	return &postsCRUDService{}
}

func (s *postsCRUDService) GetAllPosts(ctx context.Context, req *pcs.GetAllPostsRequest) (*pcs.GetAllPostsResponse, error) {
	fmt.Println(req)

	var posts []*pcs.Post

	posts = append(posts, &pcs.Post{
		Id:     11,
		UserId: 11,
		Title:  "Any title",
		Body:   "Any body",
	})

	return &pcs.GetAllPostsResponse{
		Posts: posts,
	}, nil
}

func (s *postsCRUDService) GetPost(ctx context.Context, req *pcs.GetPostRequest) (*pcs.GetPostResponse, error) {
	fmt.Println(req)

	post := pcs.Post{
		Id:     11,
		UserId: 11,
		Title:  "Any title",
		Body:   "Any body",
	}

	return &pcs.GetPostResponse{
		Post: &post,
	}, nil
}

func (s *postsCRUDService) UpdatePartialPost(ctx context.Context, req *pcs.UpdatePartialPostRequest) (*pcs.UpdatePartialPostResponse, error) {
	fmt.Println(req)

	post := pcs.Post{
		Id:     11,
		UserId: 11,
		Title:  "Any title",
		Body:   "Any body",
	}

	return &pcs.UpdatePartialPostResponse{
		UpdatedPost: &post,
	}, nil
}
func (s *postsCRUDService) DeletePost(ctx context.Context, req *pcs.DeletePostRequest) (*pcs.DeletePostResponse, error) {
	fmt.Println(req)

	post := pcs.Post{
		Id:     11,
		UserId: 11,
		Title:  "Any title",
		Body:   "Any body",
	}

	return &pcs.DeletePostResponse{
		DeletedPost: &post,
	}, nil
}
