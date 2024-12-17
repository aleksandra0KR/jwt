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

func (rp *RepositoryPostgres) Auth(refreshToken *domain.RefreshToken) error {
	user, err := rp.GetUserByGuid(refreshToken.Guid)
	if err != nil {
		log.Error(err)
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	token, err := rp.GetRefreshToken(refreshToken.Guid)
	if err != nil {
		log.Error(err)
		return err
	}

	if token != nil {
		err = rp.DeleteRefreshToken(refreshToken.Guid)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	rp.db.Create(refreshToken)
	if rp.db.Error != nil {
		log.Println(rp.db.Error)
		return rp.db.Error
	}
	return nil
}

func (rp *RepositoryPostgres) GetUserByGuid(guid string) (*domain.User, error) {
	var user domain.User
	rp.db.Where("guid = ?", guid).First(&user)
	if rp.db.Error != nil {
		log.Println(rp.db.Error)
		return nil, rp.db.Error
	}
	return &user, nil
}

func (rp *RepositoryPostgres) DeleteRefreshToken(guid string) error {
	rp.db.Where("user_guid = ?", guid).Delete(&domain.RefreshToken{})
	if rp.db.Error != nil {
		log.Println(rp.db.Error)
		return rp.db.Error
	}
	return nil
}

func (rp *RepositoryPostgres) GetRefreshToken(guid string) (*domain.RefreshToken, error) {
	var refreshToken domain.RefreshToken
	rp.db.Where("user_guid = ?", guid).First(&refreshToken)
	if rp.db.Error != nil {
		log.Println(rp.db.Error)
		return nil, rp.db.Error
	}
	return &refreshToken, nil
}
