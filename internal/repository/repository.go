package repository

import (
	"gorm.io/gorm"
	"jwt/internal/repository/postgres"
)

type Repository struct {
	RefreshTokenRepository
	UserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		RefreshTokenRepository: postgres.NewRefreshTokenRepositoryPostgres(db),
		UserRepository:         postgres.NewUserRepositoryPostgres(db),
	}
}
