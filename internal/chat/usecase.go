package chat

import (
	"context"

	"github.com/YnMann/chat_backend/internal/models"
)

const CtxUserKey = "chat"

type UseCase interface {
	GetContacts(ctx context.Context) ([]*models.Contacts, error)
}
