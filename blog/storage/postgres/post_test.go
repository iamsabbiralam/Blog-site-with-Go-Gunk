package postgres

import (
	"context"
	"gunkBlog/blog/storage"
	"testing"
)

func TestStorage_Createpost(t *testing.T) {
	
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
			got, err := s.Createpost(context.TODO(), tt.in)
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
