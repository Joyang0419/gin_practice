package item

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

func (c *CMD) BulkInsert(ctx context.Context, items []*obj.Item) (err error) {
	mItems := make([]mysql.Item, 0, len(items))
	for _, item := range items {
		mItems = append(mItems, mysql.Item{
			UserID:   item.UserID,
			Name:     item.Name,
			Category: item.Category,
			Type:     item.Type,
		})
	}

	// 使用 GORM 的 Create 方法進行批量插入
	if result := c.db.WithContext(ctx).Create(mItems); result.Error != nil {
		return fmt.Errorf("[item][CMD][BulkInsert]Create err: %w", result.Error)
	}

	return nil
}
