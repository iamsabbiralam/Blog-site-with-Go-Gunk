package category

import (
	"context"

	"gunkBlog/blog/storage"
	tcb "gunkBlog/gunk/v1/category"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Svc) Create(ctx context.Context, req *tcb.CreateCategoryRequest) (*tcb.CreateCategoryResponse, error) {
	// need to validation request
	category := storage.Category{
		ID: req.GetCategory().ID,
		CategoryName: req.GetCategory().CategoryName,
		IsCompleted: req.GetCategory().IsCompleted,
	}
	id, err := s.core.Create(context.Background(), category)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create category")
	}
	return &tcb.CreateCategoryResponse{
		ID: id,
	}, nil
}