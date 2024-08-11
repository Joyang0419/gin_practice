package item

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"gin_practice/model/mysql"
	"gin_practice/service/obj"
)

var (
	ErrInvalidUsernames = errors.New("invalid usernames")
)

type Query struct {
	db *gorm.DB
}

func NewQuery(db *gorm.DB) *Query {
	return &Query{
		db: db,
	}
}

func (q *Query) Items(ctx context.Context, filter FilterOfItems) (items []*obj.Item, err error) {
	var opt []mysql.FilterOpt
	if filter.Type != "" {
		opt = append(opt, mysql.WithType(filter.Type))
	}
	if len(filter.Usernames) > 0 {
		isValid, users, errIsValid := q.isValidUsernames(ctx, filter.Usernames)
		if errIsValid != nil {
			return nil, fmt.Errorf("[item][Query][Items]isValidUsernames err: %w", errIsValid)
		}
		if !isValid {
			return nil, fmt.Errorf("[item][Query][Items]%w", ErrInvalidUsernames)
		}
		userIDs := make([]uint64, 0, len(users))
		for _, u := range users {
			userIDs = append(userIDs, u.ID)
		}
		opt = append(opt, mysql.WithUserIDs(userIDs))
	}

	mItems, errMItems := new(mysql.Item).Items(ctx, q.db, opt...)
	if errMItems != nil {
		return nil, fmt.Errorf("[item][Query][Items]Items err: %w", errMItems)
	}

	items = make([]*obj.Item, 0, len(mItems))
	for _, mItem := range mItems {
		items = append(items, &obj.Item{
			ID:       mItem.ID,
			UserID:   mItem.UserID,
			Name:     mItem.Name,
			Category: mItem.Category,
			Type:     mItem.Type,
		})
	}

	return items, nil
}

func (q *Query) isValidUsernames(ctx context.Context, usernames []string) (bool, []*obj.User, error) {
	if len(usernames) == 0 {
		return false, nil, errors.New("[item][Query][isValidUsernames]empty usernames")
	}

	usernamesSet := make(map[string]struct{})
	for _, username := range usernames {
		usernamesSet[username] = struct{}{}
	}

	uniqueUsernames := make([]string, 0, len(usernamesSet))
	for username := range usernamesSet {
		uniqueUsernames = append(uniqueUsernames, username)
	}

	users, err := new(mysql.User).Users(ctx, q.db, mysql.WithUsernames(uniqueUsernames))
	if err != nil {
		return false, nil, fmt.Errorf("[item][Query][isValidUsernames]Users err: %w", err)
	}

	objUsers := make([]*obj.User, 0, len(users))
	for _, u := range users {
		objUsers = append(objUsers, &obj.User{
			ID:       u.ID,
			Username: u.Username,
			Password: u.Password,
		})
	}
	if len(objUsers) != len(uniqueUsernames) {
		return false, objUsers, nil
	}

	return true, objUsers, nil
}
