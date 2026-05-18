package service

import (
	"context"

	"github.com/francisjdev/community-blog/internal/database"
)

type UserStore interface {
	CreateUser(ctx context.Context, arg database.CreateUserParams) (database.User, error)
	GetUserByEmail(ctx context.Context, email string) (database.User, error)
}
