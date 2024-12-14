package postgres

import (
	"gorm.io/gorm"
)

type RefreshTokenRepositoryPostgres struct {
	db *gorm.DB
}

func NewRefreshTokenRepositoryPostgres(db *gorm.DB) *RefreshTokenRepositoryPostgres {
	return &RefreshTokenRepositoryPostgres{db: db}
}
