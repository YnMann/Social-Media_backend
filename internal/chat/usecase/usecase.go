package usecase

import (
	"context"

	"github.com/YnMann/chat_backend/internal/chat"
	"github.com/YnMann/chat_backend/internal/models"
)

type ChatUseCase struct {
	userRepo chat.ChatRepository
}

func NewChatUseCase(
	userRepo chat.ChatRepository,
) *ChatUseCase {
	return &ChatUseCase{
		userRepo: userRepo,
	}
}

func (uc *ChatUseCase) GetContacts(ctx context.Context) ([]*models.Contacts, error) {
	contacts, err := uc.chatRepo.GetContacts(ctx)
	if err != nil {
		return nil, err
	}
	return contacts, nil
}
