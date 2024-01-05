package usecase

import (
	"github.com/YnMann/chat_backend/internal/chat"
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
