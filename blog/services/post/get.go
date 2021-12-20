package post

import (
	"context"
	tpb "gunkBlog/gunk/v1/post"
)

func(s *Svc) Get(ctx context.Context, req *tpb.GetPostRequest) (*tpb.GetPostResponse, error) {
	post, err := s.postCore.GetPost(context.Background(), req.GetID())
	if err != nil {
		return nil, err
	}
	return &tpb.GetPostResponse{
		Post: &tpb.Post{
			ID:          post.ID,
			CategoryID:  post.CategoryID,
			Title:       post.Title,
			Description: post.Description,
			Image:       post.Image,
		},
	}, nil
}