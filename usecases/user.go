package usecases

import (
	"context"

	"github.com/abrahammegantoro/quickcare-bpjs-be/models"
	"github.com/abrahammegantoro/quickcare-bpjs-be/repositories"
)

type UserUsecase struct {
	repo repositories.UserRepository
}

func NewUserUsecase(repo repositories.UserRepository) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (u *UserUsecase) RegisterUser(ctx context.Context, user *models.User) (*models.User, error) {
	return u.repo.RegisterUser(ctx, user)
}