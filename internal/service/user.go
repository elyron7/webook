package service

import (
	"context"
	"errors"

	"github.com/elyron7/webook/internal/domain"
	"github.com/elyron7/webook/internal/repository"
)

var (
	ErrEmailAlreadyExists = repository.ErrEmailAlreadyExists
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) SignUp(ctx context.Context, user domain.User) error {
	if err := user.ValidateEmailAndPassword(); err != nil {
		return err
	}

	if err := user.GenerateFromPassword(); err != nil {
		return err
	}

	if err := svc.repo.Create(ctx, user); err != nil {
		if errors.Is(err, ErrEmailAlreadyExists) {
			return ErrEmailAlreadyExists
		}
		return err
	}

	return nil
}
