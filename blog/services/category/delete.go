package category

import (
	"context"
	tcb "gunkBlog/gunk/v1/category"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Svc) Delete(ctx context.Context, req *tcb.DeleteCategoryRequest) (*tcb.DeleteCategoryResponse, error) {
	if err := s.core.Delete(context.Background(), req.GetID()); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete category")
	}
	return &tcb.DeleteCategoryResponse{}, nil
}