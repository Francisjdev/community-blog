package service

import (
	"context"

	"github.com/francisjdev/community-blog/internal/database"
)

type UserStore interface {
	CreateUser(ctx context.Context, arg database.CreateUserParams) (database.User, error)
	GetUserByEmail(ctx context.Context, email string) (database.User, error)
	DeleteUserByEmail(ctx context.Context, email string) error
	CreateRefreshToken(ctx context.Context, arg database.CreateRefreshTokenParams) (database.RefreshToken, error)
}

type PostStore interface {
	CreatePost(ctx context.Context, arg database.CreatePostParams) (database.Post, error)
}
