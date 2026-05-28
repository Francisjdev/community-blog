package service

import (
	"context"

	"github.com/francisjdev/community-blog/internal/database"
)

func (s *PostService) RegisterPost(ctx context.Context, arg database.CreatePostParams) (database.Post, error) {

	newPost, err := s.store.CreatePost(ctx, arg)
	if err != nil {
		return database.Post{}, err
	}
	return newPost, nil
}
