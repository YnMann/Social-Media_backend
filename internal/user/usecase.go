package user

import (
	"context"

	"github.com/YnMann/chat_backend/internal/models"
)

const CtxUserKey = "user"

type UseCase interface {
	GetUserProfile(ctx context.Context, uID string) (*models.User, error)
}
