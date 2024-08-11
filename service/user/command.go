package user

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"gin_practice/model/mysql"
	"gin_practice/service/obj"
)

type CMD struct {
	db *gorm.DB
}

func NewCMD(db *gorm.DB) *CMD {
	return &CMD{
		db: db,
	}
}

func (C *CMD) CreateUser(ctx context.Context, username, password string) (user *obj.User, err error) {
	u := mysql.User{
		Username: username,
		Password: password,
	}

	if err = C.db.WithContext(ctx).Create(&u).Error; err != nil {
		return nil, fmt.Errorf("[repo][user][CreateUser]Create err: %w", err)
	}

	return &obj.User{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
	}, nil
}
