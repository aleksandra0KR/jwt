package domain

type RefreshToken struct {
	ID           int    `json:"id" gorm:"column:id;primaryKey"`
	Guid         string `json:"guid" gorm:"column:user_guid"`
	RefreshToken string `json:"refresh_token" gorm:"column:refresh_token"`
}
