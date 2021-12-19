package category

import (
	"context"
	"gunkBlog/blog/storage"
)

type categoryStore interface {
	Create(context.Context, storage.Category) (int64, error)
}

type CoreSvc struct {
	store	categoryStore
}

func NewCoreSvc(s categoryStore) *CoreSvc {
	return &CoreSvc{
		store: s,
	}
}

func (cs CoreSvc) Create(ctx context.Context, t storage.Category) (int64, error) {
	return cs.store.Create(ctx, t)
}