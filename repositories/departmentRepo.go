package repositories

import (
	"camvan/models"

	"gorm.io/gorm"
)

type DepartmentRepository interface {
	Create(department *models.Department) error
	FindByID(id uint) (models.Department, error)
	FindAll() ([]models.Department, error)
	Update(department *models.Department) error
	Delete(id uint) error
}

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepository{db}
}

func (r *departmentRepository) Create(department *models.Department) error {
	return r.db.Create(department).Error
}

func (r *departmentRepository) FindByID(id uint) (models.Department, error) {
	var department models.Department
	err := r.db.Preload("SubDeps").First(&department, id).Error
	return department, err
}

func (r *departmentRepository) FindAll() ([]models.Department, error) {
	var departments []models.Department
	err := r.db.Preload("SubDeps").Find(&departments).Error
	return departments, err
}

func (r *departmentRepository) Update(department *models.Department) error {
	return r.db.Save(department).Error
}

func (r *departmentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Department{}, id).Error
}
