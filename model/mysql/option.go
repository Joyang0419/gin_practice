package mysql

import (
	"gorm.io/gorm"
)

type FilterOpt func(*gorm.DB) *gorm.DB

func WithUsernames(usernames []string) FilterOpt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("username IN ?", usernames)
	}
}

func WithIDs(ids []uint64) FilterOpt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id IN ?", ids)
	}
}

func WithPasswords(passwords []string) FilterOpt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("password IN ?", passwords)
	}
}

func WithUserIDs(userIDs []uint64) FilterOpt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id IN ?", userIDs)
	}
}

func WithType(tp string) FilterOpt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("type = ?", tp)
	}
}
