package user

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"gin_practice/model/mysql"
	"gin_practice/service/obj"
)

type Query struct {
	db *gorm.DB
}

func NewQuery(db *gorm.DB) *Query {
	return &Query{
		db: db,
	}
}

func (q *Query) User(ctx context.Context, filter FilterOfUser) (user *obj.User, err error) {
	var opt []mysql.FilterOpt
	if filter.ID != nil {
		opt = append(opt, mysql.WithIDs([]uint64{*filter.ID}))
	}
	if filter.Username != nil {
		opt = append(opt, mysql.WithUsernames([]string{*filter.Username}))
	}
	if filter.Password != nil {
		opt = append(opt, mysql.WithPasswords([]string{*filter.Password}))
	}

	users, err := new(mysql.User).Users(ctx, q.db, opt...)
	if err != nil {
		return nil, fmt.Errorf("[service][user][User]Users err: %w", err)
	}
	if len(users) == 0 {
		return nil, nil
	}

	return &obj.User{
		ID:       users[0].ID,
		Username: users[0].Username,
		Password: users[0].Password,
	}, nil
}
