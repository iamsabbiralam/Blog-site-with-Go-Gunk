package postgres

import (
	"context"
	"gunkBlog/blog/storage"
)

const insertcategory = `INSERT INTO categories (category_name) VALUES (:category_name) RETURNING id;`

func (s *Storage) Create(ctx context.Context, t storage.Category) (int64, error) {
	stmt, err := s.db.PrepareNamed(insertcategory)
	if err != nil {
		return 0, nil
	}
	var id int64
	if err := stmt.Get(&id, t); err != nil {
		return 0, nil
	}
	return id, nil
}