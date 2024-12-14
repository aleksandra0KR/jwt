package postgres

import (
	"gorm.io/gorm"
	"jwt/domain"
)

type UserRepositoryPostgres struct {
	db *gorm.DB
}

func NewUserRepositoryPostgres(db *gorm.DB) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{db: db}
}

func (userRepository *UserRepositoryPostgres) Register(user *domain.User) error {
	return nil
}

func (userRepository *UserRepositoryPostgres) Auth(user *domain.User) (error, *domain.User) {
	return nil, nil
}
