package user

import (
	"context"

	"github.com/YnMann/chat_backend/internal/models"
)

type UserRepository interface {
	GetUserProfile(ctx context.Context, uID string) (*models.User, error)
}
