package postgres

import (
	"context"
	"gunkBlog/blog/storage"
)

const insertCategory = `INSERT INTO categories (category_name) VALUES (:category_name) RETURNING id;`

func (s *Storage) Create(ctx context.Context, t storage.Category) (int64, error) {
	stmt, err := s.db.PrepareNamed(insertCategory)
	if err != nil {
		return 0, err
	}
	var id int64
	if err := stmt.Get(&id, t); err != nil {
		return 0, err
	}
	// log.Println("Category ID: ", id)
	return id, nil
}

func (s *Storage) Get(ctx context.Context, id int64) (storage.Category, error) {
	var t storage.Category
	if err := s.db.Get(&t, "SELECT * FROM categories WHERE id = $1", id); err != nil {
		return t, err
	}
	return t, nil
}

const updateCategory = `UPDATE categories SET category_name = :category_name WHERE id = :id RETURNING *;`

func (s *Storage) Update(ctx context.Context, t storage.Category) error {
	stmt, err := s.db.PrepareNamed(updateCategory)
	if err != nil {
		return err
	}
	var cat storage.Category
	if err := stmt.Get(&cat, t); err != nil {
		return err
	}
	return nil
}

func (s *Storage) Delete(ctx context.Context, id int64) error {
	var t storage.Category
	if err := s.db.Get(&t, "DELETE FROM categories WHERE id = $1 RETURNING *", id); err != nil {
		return err
	}
	return nil
}

func (s *Storage) Show(ctx context.Context) ([]storage.Category, error) {
	var cat []storage.Category
	if err := s.db.Select(&cat, "SELECT * FROM categories"); err != nil {
		return cat, err
	}
	return cat, nil
}