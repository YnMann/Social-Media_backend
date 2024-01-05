package usecase

import (
	"context"

	"github.com/YnMann/chat_backend/internal/models"
	"github.com/YnMann/chat_backend/internal/user"
)

type UserUseCase struct {
	userRepo user.UserRepository
}

func NewUserUseCase(
	userRepo user.UserRepository,
) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (uc *UserUseCase) GetUserProfile(ctx context.Context, uID string) (*models.User, error) {
	c, err := uc.userRepo.GetUserProfile(ctx, uID)
	if err != nil {
		return nil, err
	}
	return c, nil
}
