package mysql

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Item struct {
	ID       uint64 `gorm:"column:id"`
	UserID   uint64 `gorm:"column:user_id"`
	Name     string `gorm:"column:name"`
	Category string `gorm:"column:category"`
	Type     string `gorm:"column:type"` // purchase
}

// TableName specifies the table name for the User model
func (*Item) TableName() string {
	return "items"
}

func (m *Item) Items(ctx context.Context, db *gorm.DB, optFilter ...FilterOpt) (items []Item, err error) {
	if db == nil {
		return nil, errors.New("[model][Item][Items]gorm is nil")
	}

	query := db.WithContext(ctx)
	for _, opt := range optFilter {
		query = opt(query)
	}

	// 執行查詢
	result := query.Find(&items)
	if resultErr := result.Error; resultErr != nil {
		if errors.Is(resultErr, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("[model][Item][Items]query.Find err: %w", resultErr)
	}

	return items, nil
}
