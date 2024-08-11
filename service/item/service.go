package item

import (
	"context"

	"gin_practice/service/obj"
)

type Service struct {
	Query IQuery
	CMD   ICommand
}

func NewService(query IQuery, CMD ICommand) *Service {
	return &Service{Query: query, CMD: CMD}
}

func (s *Service) BulkInsert(ctx context.Context, items []*obj.Item) (err error) {
	return s.CMD.BulkInsert(ctx, items)
}

func (s *Service) Items(ctx context.Context, filter FilterOfItems) (items []*obj.Item, err error) {
	return s.Query.Items(ctx, filter)
}
