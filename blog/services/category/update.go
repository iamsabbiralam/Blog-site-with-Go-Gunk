package category

import (
	"context"
	"fmt"
	"gunkBlog/blog/storage"
	tcb "gunkBlog/gunk/v1/category"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func(s *Svc) Update(ctx context.Context, req *tcb.UpdateCategoryRequest) (*tcb.UpdateCategoryResponse, error) {
	category := storage.Category{
		ID:           req.GetCategory().ID,
		CategoryName: req.GetCategory().CategoryName,
		IsCompleted:  req.GetCategory().IsCompleted,
	}
	fmt.Printf("%v", category)
	err := s.core.Update(context.Background(), category)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update category")
	}
	return &tcb.UpdateCategoryResponse{}, nil
}