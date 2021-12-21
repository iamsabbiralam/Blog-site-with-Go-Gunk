package postgres

import (
	"context"
	"gunkBlog/blog/storage"
	"log"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestShowPost(t *testing.T) {

	s := newTestStorage(t)

	tests := []struct {
		name    string
		want    []storage.Post
		wantErr bool
	}{
		{
			name: "GET_POST_LIST_SUCCESS",
			want: []storage.Post{
				{
					ID: 1,
					CategoryID: 1,
					Title: "This is title",
					Description: "This is description",
					Image: "1.jpg",
					CategoryName: "This is category update",
				},
				{
					ID: 2,
					CategoryID: 1,
					Title: "This is title 2",
					Description: "This is description 2",
					Image: "2.jpg",
					CategoryName: "This is category update",
				},
				{
					ID: 3,
					CategoryID: 1,
					Title: "This is title 23",
					Description: "This is description 3",
					Image: "3.jpg",
					CategoryName: "This is category update",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := s.ShowPost(context.Background())
			log.Printf("=========: %#v", gotList)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			sort.Slice(tt.want, func(i, j int) bool {
				return tt.want[i].ID < tt.want[j].ID
			})
			sort.Slice(gotList, func(i, j int) bool {
				return gotList[i].ID < gotList[j].ID
			})
			for i, got := range gotList {

				if !cmp.Equal(got, tt.want[i]) {
					t.Errorf("Diff: got -, want += %v", cmp.Diff(got, tt.want[i]))
				}

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

func TestDeletePost(t *testing.T) {
	
	s := newTestStorage(t)

	tests := []struct {
		name    string
		in		int64
		wantErr bool
	}{
		{
			name: "DELETE_POST_SUCCESS",
			in: 1,
		},
		{
			name: "FAILED_TO_DELETE_POST_ID",
			in: 100,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := s.DeletePost(context.TODO(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Getpost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if got != tt.want {
			// 	t.Errorf("Storage.Getpost() = %v, want %v", got, tt.want)
			// }
		})
	}
}