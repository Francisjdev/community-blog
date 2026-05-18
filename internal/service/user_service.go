package service

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/francisjdev/community-blog/internal/auth"
	"github.com/francisjdev/community-blog/internal/database"
	"github.com/francisjdev/community-blog/internal/helper"
)

func (s *UserService) RegisterUser(ctx context.Context, arg database.CreateUserParams) (database.User, error) {
	existingUser, err := s.store.GetUserByEmail(ctx, arg.Email)
	if err == nil {
		log.Printf("error: %v, user: %v\n", helper.ErrUserAlreadyExists, existingUser.Email)
		return database.User{}, helper.ErrUserAlreadyExists
	}
	if errors.Is(err, sql.ErrNoRows) {
		hashedPass, err := auth.HashPassword(arg.PasswordHash)
		if err != nil {
			log.Printf("error %v\n", err)
			return database.User{}, err
		}
		arg.PasswordHash = hashedPass
		newUser, err := s.store.CreateUser(ctx, arg)
		if err != nil {
			log.Printf("error %v\n", err)
			return database.User{}, err
		}
		return newUser, nil
	}
	log.Printf("error %v\n", err)
	return database.User{}, err

}
