package repositories

import (
	"camvan/models"

	"gorm.io/gorm"
)

type SubDepartmentRepository interface {
	Create(subDepartment *models.SubDepartment) error
	FindByID(id uint) (models.SubDepartment, error)
	FindAll() ([]models.SubDepartment, error)
	FindByDepartmentID(departmentID uint) ([]models.SubDepartment, error)
	Update(subDepartment *models.SubDepartment) error
	Delete(id uint) error
}

type subDepartmentRepository struct {
	db *gorm.DB
}

func NewSubDepartmentRepository(db *gorm.DB) SubDepartmentRepository {
	return &subDepartmentRepository{db}
}

func (r *subDepartmentRepository) Create(subDepartment *models.SubDepartment) error {
	return r.db.Create(subDepartment).Error
}

func (r *subDepartmentRepository) FindByID(id uint) (models.SubDepartment, error) {
	var subDepartment models.SubDepartment
	err := r.db.Preload("Employees").First(&subDepartment, id).Error
	return subDepartment, err
}

func (r *subDepartmentRepository) FindAll() ([]models.SubDepartment, error) {
	var subDepartments []models.SubDepartment
	err := r.db.Find(&subDepartments).Error
	return subDepartments, err
}

func (r *subDepartmentRepository) FindByDepartmentID(departmentID uint) ([]models.SubDepartment, error) {
	var subDepartments []models.SubDepartment
	err := r.db.Where("department_id = ?", departmentID).Find(&subDepartments).Error
	return subDepartments, err
}

func (r *subDepartmentRepository) Update(subDepartment *models.SubDepartment) error {
	return r.db.Save(subDepartment).Error
}

func (r *subDepartmentRepository) Delete(id uint) error {
	return r.db.Delete(&models.SubDepartment{}, id).Error
}
