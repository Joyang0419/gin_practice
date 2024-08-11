package item

import (
	"context"

	"gin_practice/service/obj"
)

type IService interface {
	BulkInsert(ctx context.Context, items []*obj.Item) (err error)
	Items(ctx context.Context, filter FilterOfItems) (items []*obj.Item, err error)
}

type FilterOfItems struct {
	Usernames []string
	Type      string // purchase
}

type IQuery interface {
	Items(ctx context.Context, filter FilterOfItems) (items []*obj.Item, err error)
	isValidUsernames(ctx context.Context, usernames []string) (bool, []*obj.User, error)
}

type ICommand interface {
	BulkInsert(ctx context.Context, items []*obj.Item) (err error)
}
