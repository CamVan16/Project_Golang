package models

type SubDepartment struct {
	ID           uint       `gorm:"primaryKey"`
	Name         string     `json:"name"`
	DepartmentID uint       `json:"department_id"`
	Employees    []Employee `gorm:"foreignKey:SubDepartmentID" json:"employees"`
}
