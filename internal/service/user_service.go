package service

import (
	"context"
	"github.com/edvina16/orbitos-backend/internal/database"
	"github.com/edvina16/orbitos-backend/internal/models"
)

type UserService struct {
	DB *database.Queries
}

func (s *UserService) CreateUser(ctx context.Context, username, email, passwordHash string) (models.User, error) {
	params := database.CreateUserParams{
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
	}
	dbUser, err := s.DB.CreateUser(ctx, params)
	if err != nil {
		return models.User{}, err
	}
	return models.User{
		ID:       int(dbUser.ID),
		Username: dbUser.Username,
		Email:    dbUser.Email,
		Password: dbUser.PasswordHash,
	}, nil
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	dbUser, err := s.DB.GetUserByUsername(ctx, username)
	if err != nil {
		return models.User{}, err
	}
	return models.User{
		ID:       int(dbUser.ID),
		Username: dbUser.Username,
		Email:    dbUser.Email,
		Password: dbUser.PasswordHash,
	}, nil
}
