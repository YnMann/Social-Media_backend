package auth

import (
	"context"

	"github.com/YnMann/chat_backend/internal/models"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, username, password string) (*models.User, error)
}
