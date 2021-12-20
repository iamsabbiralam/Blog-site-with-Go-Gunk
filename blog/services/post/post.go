package post

import (
	"context"
	"gunkBlog/blog/storage"
	tpb "gunkBlog/gunk/v1/post"
)

type postCoreStore interface {
	Createpost(context.Context, storage.Post) (int64, error)
}

type Svc struct {
	tpb.UnimplementedPostServiceServer
	postCore	postCoreStore
}

func NewPostServer() * Svc {
	return &Svc{}
}