package postgres

import (
	"context"
	"gunkBlog/blog/storage"
	"testing"
)

func TestCreatePost(t *testing.T) {
	
	s := newTestStorage(t)

	tests := []struct {
		name    string
		in		storage.Post
		want    int64
		wantErr bool
	}{
		{
			name: "CREATE_POST_SUCCESS",
			in: storage.Post{
				ID:          1,
				CategoryID:  1,
				Title:       "This is title",
				Description: "This is description",
				Image:       "1.jpg",
			},
			want: 1,
		},
		{
			name: "FAILED_POST_TITLE_DUPLICATE",
			in: storage.Post{
				ID:          1,
				CategoryID:  1,
				Title:       "This is title",
				Description: "This is description",
				Image:       "1.jpg",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreatePost(context.TODO(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Createpost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.Createpost() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestGetPost(t *testing.T) {
	
	s := newTestStorage(t)

	tests := []struct {
		name    string
		in		int64
		want    storage.Post
		wantErr bool
	}{
		{
			name: "GET_POST_SUCCESS",
			in: 1,
			want: storage.Post{
				ID:          1,
				CategoryID:  1,
				Title:       "This is title",
				Description: "This is description",
				Image:       "1.jpg",
			},
		},
		{
			name: "FAILED_TO_GET_POST_ID",
			in: 100,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetPost(context.TODO(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Getpost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.Getpost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdatePost(t *testing.T) {
	
	s := newTestStorage(t)

	tests := []struct {
		name    string
		in		storage.Post
		want    *storage.Post
		wantErr bool
	}{
		{
			name: "UPDATE_POST_SUCCESS",
			in: storage.Post{
				ID:          1,
				CategoryID:  1,
				Title:       "This is title update",
				Description: "This is description update",
				Image:       "2.jpg",
			},
			want: &storage.Post{
				ID:          1,
				CategoryID:  1,
				Title:       "This is title update",
				Description: "This is description update",
				Image:       "2.jpg",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := s.UpdatePost(context.TODO(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Updatepost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if got != tt.want {
			// 	t.Errorf("Storage.Getpost() = %v, want %v", got, tt.want)
			// }
		})
	}
}
