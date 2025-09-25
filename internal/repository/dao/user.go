package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
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
	user.CreatedAt = time.Now().UnixMilli()
	user.UpdatedAt = time.Now().UnixMilli()

	err := dao.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return ErrEmailAlreadyExists
		}
		return err
	}
	return nil
}

type User struct {
	Id        uint64 `gorm:"primary_key,AUTO_INCREMENT"`
	Email     string `gorm:"type:varchar(255);uniqueIndex:idx_email"`
	Password  string `gorm:"type:varchar(255)"`
	CreatedAt int64
	UpdatedAt int64
}
