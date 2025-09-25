package domain

import (
	"errors"

	"github.com/dlclark/regexp2"
	"golang.org/x/crypto/bcrypt"
)

const (
	emailRegex    = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	passwordRegex = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`
)

var (
	ErrSystemFailure         = errors.New("system error")
	ErrInvalidEmail          = errors.New("invalid email")
	ErrInvalidPassword       = errors.New("invalid password")
	ErrPasswordMismatch      = errors.New("password mismatch")
	ErrPasswordHashingFailed = errors.New("password hashing failed")
	ErrPasswordCompareFailed = errors.New("password comparison failed")
)

type User struct {
	Email           string
	Password        string
	ConfirmPassword string
}

func (user *User) ValidateEmailAndPassword() error {
	ok, err := user.ValidateEmail()
	if err != nil {
		return ErrSystemFailure
	}
	if !ok {
		return ErrInvalidEmail
	}

	ok, err = user.ValidatePassword()
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

func (user *User) ValidatePassword() (bool, error) {
	re := regexp2.MustCompile(passwordRegex, 0)
	return re.MatchString(user.Password)
}

func (user *User) ValidateEmail() (bool, error) {
	re := regexp2.MustCompile(emailRegex, 0)
	return re.MatchString(user.Email)
}

// GenerateFromPassword generates a password hash
func (user *User) GenerateFromPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return ErrPasswordHashingFailed
	}
	user.Password = string(hash)
	return nil
}

// ComparePasswords compares the password with the stored hash
func (user *User) ComparePasswords(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return ErrPasswordCompareFailed
	}
	return nil
}
