package repository

import (
	"context"

	"github.com/elyron7/webook/internal/domain"
	"github.com/elyron7/webook/internal/repository/dao"
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
		return err
	}

	return nil
}
