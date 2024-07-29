package models

type Employee struct {
	IDEm            uint   `gorm:"primaryKey" json:"idem"`
	Name            string `json:"name"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	SubDepartmentID uint   `json:"sub_department_id"`
	// AccessToken     string `json:"accessToken"`
	// RefreshToken    string `json:"refreshToken"`
}
