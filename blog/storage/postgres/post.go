package postgres

import (
	"context"
	"gunkBlog/blog/storage"
)

const insertPost = `INSERT INTO posts (cat_id, title, description, image) VALUES (:cat_id, :title, :description, :image) RETURNING id;`

func (s *Storage) CreatePost(ctx context.Context, t storage.Post) (int64, error) {
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

func(s *Storage) ShowPost(ctx context.Context) ([]storage.Post, error) {
	var post []storage.Post
	if err := s.db.Select(&post, "SELECT posts.* , categories.category_name FROM posts LEFT JOIN categories ON categories.id = posts.cat_id"); err != nil {
		return nil, err
	}
	return post, nil
}

func(s *Storage) GetPost(ctx context.Context, id int64) (storage.Post, error) {
	var p storage.Post
	if err := s.db.Get(&p, "SELECT * FROM posts WHERE id = $1", id); err != nil {
		return p, err
	}
	return p, nil
}

const updatePost = `UPDATE posts SET cat_id = :cat_id, title = :title, description = :description, image = :image WHERE id = :id RETURNING *;`

func(s *Storage) UpdatePost(ctx context.Context, t storage.Post) error {
	stmt, err := s.db.PrepareNamed(updatePost)
	if err != nil {
		return err
	}
	var p storage.Post
	if err := stmt.Get(&p, t); err != nil {
		return err
	}
	return nil
}

func(s *Storage) DeletePost(ctx context.Context, id int64) error {
	var p storage.Post
	if err := s.db.Get(&p, "DELETE FROM posts WHERE id = $1 RETURNING *", id); err != nil {
		return err
	}
	return nil
}