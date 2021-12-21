package post

import (
	"context"
	"fmt"
	"gunkBlog/blog/storage"
	"gunkBlog/blog/storage/postgres"
)

type CoreSvc struct {
	postStore	*postgres.Storage
}

func NewCoreSvc(s *postgres.Storage) *CoreSvc {
	return &CoreSvc{
		postStore: s,
	}
}

func(cs CoreSvc) CreatePost(ctx context.Context, t storage.Post) (int64, error) {
	fmt.Println(t, ctx)
	return cs.postStore.CreatePost(ctx, t)
}

func(cs CoreSvc) ShowPost(ctx context.Context) ([]storage.Post ,error) {
	return cs.postStore.ShowPost(ctx)
}

func(cs CoreSvc) GetPost(ctx context.Context, id int64) (storage.Post, error) {
	return cs.postStore.GetPost(ctx, id)
}

func(cs CoreSvc) UpdatePost(ctx context.Context, t storage.Post) error {
	return cs.postStore.UpdatePost(ctx, t)
}

func(cs CoreSvc) DeletePost(ctx context.Context, id int64) error {
	return cs.postStore.DeletePost(ctx, id)
}