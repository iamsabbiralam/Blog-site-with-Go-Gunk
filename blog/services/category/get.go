package category

import (
	"context"
	"gunkBlog/blog/storage"
	tcb "gunkBlog/gunk/v1/category"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func(s *Svc) Get(ctx context.Context, req *tcb.GetCategoryRequest) (*tcb.GetCategoryResponse, error) {
	var cat storage.Category
	cat, err := s.core.Get(context.Background(), req.GetID())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get category")
	}
	return &tcb.GetCategoryResponse{
		Category: &tcb.Category{
			ID:           cat.ID,
			CategoryName: cat.CategoryName,
			IsCompleted:  cat.IsCompleted,
		},
	}, nil
}