package postgres

import (
	"context"
	"gunkBlog/blog/storage"
)

const insertPost = `INSERT INTO posts (cat_id, title, description, image) VALUES (:cat_id, :title, :description, :image) RETURNING id;`

func (s *Storage) Createpost(ctx context.Context, t storage.Post) (int64, error) {
	stmt, err := s.db.PrepareNamed(insertPost)
	if err != nil {
		return 0, err
	}
	var id int64
	if err := stmt.Get(&id, t); err != nil {
		return 0, err
	}
	return id, nil
}