package usecase

import (
	"context"

	"github.com/YnMann/chat_backend/internal/chat"
	"github.com/YnMann/chat_backend/internal/models"
)

type ChatUseCase struct {
	chatRepo chat.ChatRepository
}

func NewChatUseCase(
	chatRepo chat.ChatRepository,
) *ChatUseCase {
	return &ChatUseCase{
		chatRepo: chatRepo,
	}
}

func (uc *ChatUseCase) GetContacts(ctx context.Context) ([]*models.Contacts, error) {
	c, err := uc.chatRepo.GetContacts(ctx)
	if err != nil {
		return nil, err
	}
	return c, nil
}
