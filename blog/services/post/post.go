package post

import (
	"context"
	"gunkBlog/blog/storage"
	tpb "gunkBlog/gunk/v1/post"
)

type postCoreStore interface {
	CreatePost(context.Context, storage.Post) (int64, error)
	ShowPost(context.Context) ([]storage.Post ,error)
	GetPost(context.Context, int64) (storage.Post, error)
	UpdatePost(context.Context, storage.Post) error
	DeletePost(context.Context, int64) error
}

type Svc struct {
	tpb.UnimplementedPostServiceServer
	postCore	postCoreStore
}

func NewPostServer(p postCoreStore) * Svc {
	return &Svc{
		postCore: p,
	}
}