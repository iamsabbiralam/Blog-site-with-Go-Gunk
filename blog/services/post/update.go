package post

import (
	"context"
	"gunkBlog/blog/storage"
	tpb "gunkBlog/gunk/v1/post"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func(s *Svc) Update(ctx context.Context, req *tpb.UpdatePostRequest) (*tpb.UpdatePostResponse, error) {
	post := storage.Post{
		ID:          req.Post.ID,
		CategoryID:  req.Post.CategoryID,
		Title:       req.Post.Title,
		Description: req.Post.Description,
		Image:       req.Post.Title,
	}
	if err := s.postCore.UpdatePost(context.Background(), post); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update post")
	}
	return &tpb.UpdatePostResponse{}, nil
}