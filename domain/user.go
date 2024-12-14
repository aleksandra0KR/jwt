package domain

type User struct {
	Guid         string `json:"guid" gorm:"column:guid primary_key"`
	Email        string `json:"email" gorm:"column:email"`
	IP           string `json:"ip" gorm:"column:ip"`
	AccessToken  string `json:"access_token" gorm:"column:access_token"`
	RefreshToken string `json:"refresh_token" gorm:"column:refresh_token"`
}
