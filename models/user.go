package models

type User struct {
	IDUser       uint   `gorm: "primaryKey"`
	Phone        string `json: "phone"`
	Password     string `json: "password"`
	ConfirmPass  string `json: "confirmpass"`
	IsEmployee   bool   `json:"isemployee"`
	AccessToken  string `json: "accessToken`
	RefreshToken string `json: "refreshToken"`
}
