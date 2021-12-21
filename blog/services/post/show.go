package post

import (
	"context"
	tpb "gunkBlog/gunk/v1/post"
)

func(s *Svc) Show(ctx context.Context, req *tpb.ShowPostRequest) (*tpb.ShowPostResponse, error) {
	posts, err := s.postCore.ShowPost(context.Background())
	if err != nil {
		return nil, err
	}
	var p []*tpb.Post
	for _, value := range posts {
		p = append(p, &tpb.Post{
			ID:          value.ID,
			CategoryID:  value.CategoryID,
			Title:       value.Title,
			Description: value.Description,
			Image:       value.Image,
			CategoryName: value.CategoryName,
		})
	}
	return &tpb.ShowPostResponse{
		Post: p,
	}, nil
}