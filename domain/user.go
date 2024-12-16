package domain

type User struct {
	Guid         string `json:"guid" gorm:"column:guid;primaryKey"`
	Email        string `json:"email" gorm:"column:email"`
	IP           string `json:"ip" gorm:"column:ip"`
	AccessToken  string `json:"access_token" gorm:"-"`
	RefreshToken string `json:"refresh_token" gorm:"-"`
}
