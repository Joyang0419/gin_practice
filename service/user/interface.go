package user

import (
	"context"

	"gin_practice/service/obj"
)

type IService interface {
	Register(ctx context.Context, username, password string) (user *obj.User, err error)
	isUsernameExist(ctx context.Context, username string) (bool, error)
	Login(ctx context.Context, username, password string) (user *obj.User, err error)
}

type IQuery interface {
	User(ctx context.Context, filter FilterOfUser) (user *obj.User, err error)
}

type ICommand interface {
	CreateUser(ctx context.Context, username, password string) (user *obj.User, err error)
}

type FilterOfUser struct {
	ID       *uint64
	Username *string
	Password *string
}
