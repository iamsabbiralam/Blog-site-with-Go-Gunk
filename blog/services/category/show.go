package category

import (
	"context"
	tcb "gunkBlog/gunk/v1/category"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Svc) Show(ctx context.Context, req *tcb.ShowCategoryRequest) (*tcb.ShowCategoryResponse, error) {
	cats, err := s.core.Show(context.Background())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get categories: %s", err)
	}
	var c []*tcb.Category
	for _, val := range cats {
		c = append(c, &tcb.Category{
			ID:           val.ID,
			CategoryName: val.CategoryName,
			IsCompleted:  val.IsCompleted,
		})
	}
	return &tcb.ShowCategoryResponse{
		Category: c,
	}, nil
}