package services

import (
	"context"

	"github.com/dstm45/vacances/pkg/database"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserService struct {
	db *database.Queries
}

func NewUserService(db *database.Queries) *UserService {
	userService := UserService{
		db: db,
	}
	return &userService
}

func (s UserService) CreateUser(ctx context.Context, params database.CreateUserParams) (*database.User, error) {
	user, err := s.db.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := s.db.DeleteUser(ctx, id)
	return err
}

func (s UserService) GetUser(ctx context.Context, id uuid.UUID) (*database.User, error) {
	user, err := s.db.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s UserService) ChangeEmail(ctx context.Context, email string) error {
	err := s.db.ChangeEmail(ctx, pgtype.Text{String: email, Valid: true})
	return err
}

func (s UserService) ChangeUsername(ctx context.Context, username string) error {
	err := s.db.ChangeUsername(ctx, pgtype.Text{String: username, Valid: true})
	return err
}

func (s UserService) ChangePassword(ctx context.Context, passwordHash string) error {
	err := s.db.ChangeEmail(ctx, pgtype.Text{String: passwordHash, Valid: true})
	return err
}
