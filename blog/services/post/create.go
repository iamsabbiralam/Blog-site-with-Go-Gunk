package post

import (
	"context"
	"fmt"
	"gunkBlog/blog/storage"
	tpb "gunkBlog/gunk/v1/post"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func(s *Svc) Create(ctx context.Context, req *tpb.CreatePostRequest) (*tpb.CreatePostResponse, error) {
	post := storage.Post{
		CategoryID:  req.Post.CategoryID,
		Title:       req.Post.Title,
		Description: req.Post.Description,
		Image:       req.Post.Image,
	}
	id, err := s.postCore.CreatePost(context.Background(), post)

	fmt.Println(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create post")
	}
	return &tpb.CreatePostResponse{
		ID: id,
	}, nil
}