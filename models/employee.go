package models

type Employee struct {
	ID              uint   `gorm:"primaryKey"`
	Name            string `json:"name"`
	SubDepartmentID uint   `json:"sub_department_id"`
}
