package post

import (
	"context"
	tpb "gunkBlog/gunk/v1/post"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func(s *Svc) Delete(ctx context.Context, req *tpb.DeletePostRequest) (*tpb.DeletePostResponse, error) {
	if err := s.postCore.DeletePost(context.Background(), req.ID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete post")
	}
	return &tpb.DeletePostResponse{}, nil
}