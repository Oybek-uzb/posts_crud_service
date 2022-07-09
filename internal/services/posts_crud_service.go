package services

import (
	"context"
	"time"

	pbpc "github.com/Oybek-uzb/posts_crud_service/pkg/api/posts_crud_service"
	pbp "github.com/Oybek-uzb/posts_crud_service/pkg/api/posts_service"
	"github.com/Oybek-uzb/posts_crud_service/services"
)

type postsCRUDService struct {
	pbpc.UnimplementedPostsCRUDServiceServer
	Services services.ServiceManager
}

func NewPostsCRUDService(ss services.ServiceManager) *postsCRUDService {
	return &postsCRUDService{
		Services: ss,
	}
}

func (s *postsCRUDService) GetAllPosts(ctx context.Context, req *pbpc.GetAllPostsRequest) (*pbpc.GetAllPostsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := s.Services.PostsService().GetAllPosts(ctx, &pbp.GetAllPostsRequest{})
	if err != nil {
		return nil, err
	}

	var allPosts []*pbpc.Post
	var postsLen = len(response.Posts)

	if postsLen == 0 {
		return &pbpc.GetAllPostsResponse{
			Posts: []*pbpc.Post{},
		}, nil
	}
	for _, post := range response.Posts {
		newPost := pbpc.Post{
			Id:     post.Id,
			UserId: post.UserId,
			Title:  post.Title,
			Body:   post.Body,
		}

		allPosts = append(allPosts, &newPost)
	}

	return &pbpc.GetAllPostsResponse{
		Posts: allPosts,
	}, nil
}

func (s *postsCRUDService) GetPost(ctx context.Context, req *pbpc.GetPostRequest) (*pbpc.GetPostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := s.Services.PostsService().GetPost(ctx, &pbp.GetPostRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	if response.Post == nil {
		return &pbpc.GetPostResponse{
			Post: nil,
		}, nil
	}

	var newPost = &pbpc.Post{
		Id:     response.Post.Id,
		UserId: response.Post.UserId,
		Title:  response.Post.Title,
		Body:   response.Post.Body,
	}

	return &pbpc.GetPostResponse{
		Post: newPost,
	}, nil
}

func (s *postsCRUDService) UpdatePartialPost(ctx context.Context, req *pbpc.UpdatePartialPostRequest) (*pbpc.UpdatePartialPostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	var updateData = pbp.Post{
		Id:     req.UpdateData.Id,
		UserId: req.UpdateData.UserId,
		Title:  req.UpdateData.Title,
		Body:   req.UpdateData.Body,
	}

	response, err := s.Services.PostsService().UpdatePartialPost(ctx, &pbp.UpdatePartialPostRequest{
		UpdateData: &updateData,
	})
	if err != nil {
		return nil, err
	}

	if response.UpdatedPost == nil {
		return &pbpc.UpdatePartialPostResponse{
			UpdatedPost: nil,
		}, nil
	}

	var newPost = &pbpc.Post{
		Id:     response.UpdatedPost.Id,
		UserId: response.UpdatedPost.UserId,
		Title:  response.UpdatedPost.Title,
		Body:   response.UpdatedPost.Body,
	}

	return &pbpc.UpdatePartialPostResponse{
		UpdatedPost: newPost,
	}, nil
}
func (s *postsCRUDService) DeletePost(ctx context.Context, req *pbpc.DeletePostRequest) (*pbpc.DeletePostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := s.Services.PostsService().DeletePost(ctx, &pbp.DeletePostRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	if response.DeletedPost == nil {
		return &pbpc.DeletePostResponse{
			DeletedPost: nil,
		}, nil
	}

	var deletedPost = &pbpc.Post{
		Id:     response.DeletedPost.Id,
		UserId: response.DeletedPost.UserId,
		Title:  response.DeletedPost.Title,
		Body:   response.DeletedPost.Body,
	}

	return &pbpc.DeletePostResponse{
		DeletedPost: deletedPost,
	}, nil
}
