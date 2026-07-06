package service

import (
	"context"

	"github.com/francisjdev/community-blog/internal/database"
	"github.com/google/uuid"
)

func (s *PostService) RegisterPost(ctx context.Context, arg database.CreatePostParams) (database.Post, error) {

	newPost, err := s.store.CreatePost(ctx, arg)
	if err != nil {
		return database.Post{}, err
	}
	return newPost, nil
}
func (s *PostService) ListAllPostsByUser(ctx context.Context, userID uuid.UUID) ([]database.Post, error) {

	posts, err := s.store.ListAllPostsByUser(ctx, userID)
	if err != nil {
		return []database.Post{}, err
	}
	return posts, nil
}
func (s *PostService) GetPostById(ctx context.Context, id uuid.UUID) (database.Post, error) {

	post, err := s.store.GetPostById(ctx, id)
	if err != nil {
		return database.Post{}, err
	}
	return post, nil
}

func (s *PostService) DeleteSinglePost(ctx context.Context, arg database.DeletePostByIdParams) error {

	err := s.store.DeletePostById(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostService) DeleteAllPostsByUser(ctx context.Context, userID uuid.UUID) error {

	err := s.store.DeleteAllPostByUser(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostService) ListAllPosts(ctx context.Context) ([]database.Post, error) {

	posts, err := s.store.ListAllPosts(ctx)
	if err != nil {
		return []database.Post{}, err
	}
	return posts, nil
}
