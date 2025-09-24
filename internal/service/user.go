package service

import (
	"context"
	"errors"

	"github.com/dlclark/regexp2"
	"github.com/elyron7/webook/internal/domain"
	"github.com/elyron7/webook/internal/repository"
)

const (
	emailRegex    = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	passwordRegex = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`
)

var (
	ErrSystemFailure    = errors.New("system error")
	ErrInvalidEmail     = errors.New("invalid email")
	ErrInvalidPassword  = errors.New("invalid password")
	ErrPasswordMismatch = errors.New("password mismatch")
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
	if err := svc.ValidateEmailAndPassword(user); err != nil {
		return err
	}

	if err := svc.repo.Create(ctx, user); err != nil {
		return err
	}
	return nil
}

func (svc *UserService) ValidateEmailAndPassword(user domain.User) error {
	ok, err := ValidateEmail(user.Email)
	if err != nil {
		return ErrSystemFailure
	}
	if !ok {
		return ErrInvalidEmail
	}

	ok, err = ValidatePassword(user.Password)
	if err != nil {
		return ErrSystemFailure
	}
	if !ok {
		return ErrInvalidPassword
	}

	if user.Password != user.ConfirmPassword {
		return ErrPasswordMismatch
	}

	return nil
}

func ValidatePassword(password string) (bool, error) {
	re := regexp2.MustCompile(passwordRegex, 0)
	return re.MatchString(password)
}

func ValidateEmail(email string) (bool, error) {
	re := regexp2.MustCompile(emailRegex, 0)
	return re.MatchString(email)
}
