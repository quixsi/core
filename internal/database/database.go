package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/quixsi/app/internal/models"
)

type Database interface {
	GetUsers(ctx context.Context) ([]*models.User, error)
	GetUser(ctx context.Context, id uuid.UUID) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}