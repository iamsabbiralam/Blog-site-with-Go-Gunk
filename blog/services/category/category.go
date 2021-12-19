package category

import (
	"context"

	"gunkBlog/blog/storage"
	tcb "gunkBlog/gunk/v1/category"
)

type categoryCoreStore interface {
	Create(context.Context, storage.Category) (int64, error)
}

type Svc struct {
	tcb.UnimplementedCategoryServiceServer
	core categoryCoreStore
}

func NewCategoryServer(c categoryCoreStore) *Svc {
	return &Svc{
		core: c,
	}
}