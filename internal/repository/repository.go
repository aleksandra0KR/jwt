package repository

import (
	"gorm.io/gorm"
	"jwt/domain"
	"jwt/internal/repository/postgres"
)

type Repository interface {
	Auth(*domain.RefreshToken) error
	GetUserByGuid(string) (*domain.User, error)
	DeleteRefreshToken(string) error
	GetRefreshToken(string) (*domain.RefreshToken, error)
}

func NewRepository(db *gorm.DB) Repository {
	return postgres.NewRepositoryPostgres(db)
}
