package service

import (
	"context"
	"errors"

	"github.com/elyron7/webook/internal/domain"
	"github.com/elyron7/webook/internal/repository"
)

var (
	ErrEmailAlreadyExists    = repository.ErrEmailAlreadyExists
	ErrSystemFailure         = domain.ErrSystemFailure
	ErrInvalidEmail          = domain.ErrInvalidEmail
	ErrPasswordCompareFailed = domain.ErrPasswordCompareFailed
	ErrUserNotFound          = repository.ErrUserNotFound
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
		if errors.Is(err, ErrSystemFailure) {
			return ErrSystemFailure
		}
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

func (svc *UserService) Login(ctx context.Context, user *domain.User) error {
	u, err := svc.repo.FindByEmail(ctx, user.Email)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return ErrUserNotFound
		}
		return nil
	}
	
	user.Id = u.Id

	err = u.ComparePasswords(user.Password)
	if err != nil {
		if errors.Is(err, ErrPasswordCompareFailed) {
			return ErrPasswordCompareFailed
		}
		return err
	}

	return nil
}
