package mysql

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// User represents the user model
type User struct {
	ID       uint64 `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

// TableName specifies the table name for the User model
func (*User) TableName() string {
	return "users"
}

func (m *User) Users(ctx context.Context, db *gorm.DB, optFilter ...FilterOpt) (users []User, err error) {
	if db == nil {
		return nil, errors.New("[model][User][Users]gorm is nil")
	}

	query := db.WithContext(ctx)
	for _, opt := range optFilter {
		query = opt(query)
	}

	// 執行查詢
	result := query.Find(&users)
	if resultErr := result.Error; resultErr != nil {
		if errors.Is(resultErr, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("[model][User][Users]query.Find err: %w", resultErr)
	}

	return users, nil
}
