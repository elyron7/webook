package dao

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) Insert(ctx context.Context, user User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	fmt.Printf("%+v \n", user)
	return dao.db.WithContext(ctx).Create(&user).Error
}

type User struct {
	Id        int64  `gorm:"primary_key,AUTO_INCREMENT"`
	Email     string `gorm:"unique_index"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
