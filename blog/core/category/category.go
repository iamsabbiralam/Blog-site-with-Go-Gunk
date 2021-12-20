package category

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

func (cs CoreSvc) Create(ctx context.Context, t storage.Category) (int64, error) {
	return cs.store.Create(ctx, t)
}

func (cs CoreSvc) Show(ctx context.Context) ([]storage.Category, error) {
	return cs.store.Show(ctx)
}

func (cs CoreSvc) Get(ctx context.Context, id int64) (storage.Category, error) {
	return cs.store.Get(ctx, id)
}

func (cs CoreSvc) Update(ctx context.Context, t storage.Category)  error {
	return cs.store.Update(ctx, t)
}