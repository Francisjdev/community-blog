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
func (s *UserService) Login(ctx context.Context, email string, password string, key string) (accessToken string, refreshToken string, error error) {
	existingUser, err := s.store.GetUserByEmail(ctx, email)
	if err != nil {
		log.Printf("error %v\n", err)
		return "", "", err
	}
	matchedPass, err := auth.CheckHash(password, existingUser.PasswordHash)
	if err != nil {
		log.Printf("error: %v\n", err)
		return "", "", err
	}
	if !matchedPass {
		log.Printf("email or password incorrect: %v\n", err)
		return "", "", helper.ErrEmailOrPasswordIncorrect
	}
	accessToken, err = auth.GenerateJWT(existingUser.ID, existingUser.Role, key)
	if err != nil {
		return "", "", err
	}
	refreshToken = auth.MakeRefreshToken()
	// here i need to insert the refresh token into the db
	hashedRefresh, err := auth.HashPassword(refreshToken)
	if err != nil {
		return "", "", err
	}
	refreshTokenParams := database.CreateRefreshTokenParams{
		UserID:    existingUser.ID,
		TokenHash: hashedRefresh,
	}

	_, err = s.store.CreateRefreshToken(ctx, refreshTokenParams)
	if err != nil {
		log.Printf("error inserting data into db: %v", err)
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *UserService) DeleteUser(ctx context.Context, email string) error {
	err := s.store.DeleteUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	return nil
}
