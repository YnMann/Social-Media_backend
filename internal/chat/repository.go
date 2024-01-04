package chat

import (
	"context"

	"github.com/YnMann/chat_backend/internal/models"
)

type ChatRepository interface {
	CreateMsg(ctx context.Context, m *models.Messages) error
	GetMsg(ctx context.Context, sender string, sender_ip string, recipient string) (*models.Messages, error)
	GetContacts(ctx context.Context) ([]*models.Contacts, error)
}
