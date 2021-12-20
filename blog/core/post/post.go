package post

import (
	"context"
	"gunkBlog/blog/storage"
	"gunkBlog/blog/storage/postgres"
)

type CoreSvc struct {
	store	*postgres.Storage
}

func NewCoreSvc(s *postgres.Storage) *CoreSvc {
	return &CoreSvc{
		store: s,
	}
}

func(cs CoreSvc) CreatePost(ctx context.Context, sp storage.Post) (int64, error) {
	return cs.store.CreatePost(ctx, sp)
}

func(cs CoreSvc) GetPost(ctx context.Context, id int64) (storage.Post, error) {
	return cs.store.GetPost(ctx, id)
}

func(cs CoreSvc) UpdatePost(ctx context.Context, t storage.Post) error {
	return cs.store.UpdatePost(ctx, t)
}