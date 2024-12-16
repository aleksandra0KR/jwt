package postgres

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"jwt/domain"
)

type RepositoryPostgres struct {
	db *gorm.DB
}

func NewRepositoryPostgres(db *gorm.DB) *RepositoryPostgres {
	return &RepositoryPostgres{db: db}
}

func (repositoryPostgres *RepositoryPostgres) Auth(refreshToken *domain.RefreshToken) error {
	user, err := repositoryPostgres.GetUserByGuid(refreshToken.Guid)
	if err != nil {
		log.Error(err)
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	token, err := repositoryPostgres.GetRefreshToken(refreshToken.Guid)
	if err != nil {
		log.Error(err)
		return err
	}

	if token != nil {
		err = repositoryPostgres.DeleteRefreshToken(refreshToken.Guid)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	repositoryPostgres.db.Create(refreshToken)
	if repositoryPostgres.db.Error != nil {
		log.Println(repositoryPostgres.db.Error)
		return repositoryPostgres.db.Error
	}
	return nil
}

func (repositoryPostgres *RepositoryPostgres) GetUserByGuid(guid string) (*domain.User, error) {
	var user domain.User
	repositoryPostgres.db.Where("guid = ?", guid).First(&user)
	if repositoryPostgres.db.Error != nil {
		log.Println(repositoryPostgres.db.Error)
		return nil, repositoryPostgres.db.Error
	}
	return &user, nil
}

func (repositoryPostgres *RepositoryPostgres) DeleteRefreshToken(guid string) error {
	repositoryPostgres.db.Where("user_guid = ?", guid).Delete(&domain.RefreshToken{})
	if repositoryPostgres.db.Error != nil {
		log.Println(repositoryPostgres.db.Error)
		return repositoryPostgres.db.Error
	}
	return nil
}

func (repositoryPostgres *RepositoryPostgres) GetRefreshToken(guid string) (*domain.RefreshToken, error) {
	var refreshToken domain.RefreshToken
	repositoryPostgres.db.Where("user_guid = ?", guid).First(&refreshToken)
	if repositoryPostgres.db.Error != nil {
		log.Println(repositoryPostgres.db.Error)
		return nil, repositoryPostgres.db.Error
	}
	return &refreshToken, nil
}
