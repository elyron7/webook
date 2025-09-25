package repository

import (
	"context"
	"errors"

	"github.com/elyron7/webook/internal/domain"
	"github.com/elyron7/webook/internal/repository/dao"
)

var (
	ErrEmailAlreadyExists = dao.ErrEmailAlreadyExists
	ErrUserNotFound       = dao.ErrUserNotFound
)

type UserRepository struct {
	userDAO *dao.UserDAO
}

func NewUserRepository(userDAO *dao.UserDAO) *UserRepository {
	return &UserRepository{
		userDAO: userDAO,
	}
}
func (repo *UserRepository) Create(ctx context.Context, user domain.User) error {
	userDAO := dao.User{
		Email:    user.Email,
		Password: user.Password,
	}

	err := repo.userDAO.Insert(ctx, userDAO)
	if err != nil {
		if errors.Is(err, ErrEmailAlreadyExists) {
			return ErrEmailAlreadyExists
		}
		return err
	}
	return nil
}

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := repo.userDAO.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return domain.User{}, ErrUserNotFound
		}
		return domain.User{}, err
	}

	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}
