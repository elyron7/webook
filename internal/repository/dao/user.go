package dao

import (
	"context"
	"log"
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
	err := dao.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		log.Println(err)
	}
	return nil
}

type User struct {
	Id        int64  `gorm:"primary_key,AUTO_INCREMENT"`
	Email     string `gorm:"unique_index"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
