package models

type Department struct {
	ID      uint            `gorm:"primaryKey"`
	Name    string          `json:"name"`
	SubDeps []SubDepartment `gorm:"foreignKey:DepartmentID" json:"sub"`
}
