package post

import (
	"context"
	"gunkBlog/blog/storage"
	tpb "gunkBlog/gunk/v1/post"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func(s *Svc) Create(ctx context.Context, req *tpb.CreatePostRequest) (*tpb.CreatePostResponse, error) {
	post := storage.Post{
		ID:          req.GetPost().ID,
		CategoryID:  req.GetPost().CategoryID,
		Title:       req.Post.Title,
		Description: req.Post.Description,
		Image:       req.Post.Image,
	}
	id, err := s.postCore.Createpost(context.Background(), post)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create post")
	}
	return &tpb.CreatePostResponse{
		ID: id,
	}, nil
}