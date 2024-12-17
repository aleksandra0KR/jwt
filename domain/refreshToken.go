package domain

type RefreshToken struct {
	ID           int    `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Guid         string `json:"guid" gorm:"column:user_guid;not null"`
	RefreshToken string `json:"refresh_token" gorm:"column:refresh_token;not null"`
	User         User   `gorm:"foreignKey:user_guid;references:guid;constraint:OnDelete:CASCADE;"`
}
