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

func(cs CoreSvc) Createpost(ctx context.Context, sp storage.Post) (int64, error) {
	return cs.store.Createpost(ctx, sp)
}