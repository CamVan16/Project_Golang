package repositories

import (
	"camvan/models"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(employee *models.Employee) error
	FindByID(id uint) (models.Employee, error)
	FindAll() ([]models.Employee, error)
	Update(employee *models.Employee) error
	Delete(id uint) error
	FindByPhonePass(phone, pass string) (models.Employee, error)
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db}
}

func (r *employeeRepository) Create(employee *models.Employee) error {
	return r.db.Create(employee).Error
}

func (r *employeeRepository) FindByID(id uint) (models.Employee, error) {
	var employee models.Employee
	err := r.db.First(&employee, id).Error
	return employee, err
}

func (r *employeeRepository) FindAll() ([]models.Employee, error) {
	var employees []models.Employee
	err := r.db.Find(&employees).Error
	return employees, err
}

func (r *employeeRepository) Update(employee *models.Employee) error {
	return r.db.Save(employee).Error
}

func (r *employeeRepository) Delete(id uint) error {
	return r.db.Delete(&models.Employee{}, id).Error
}

func (r *employeeRepository) FindByPhonePass(phone, pass string) (models.Employee, error) {
	var employee models.Employee
	err := r.db.Where("Phone = ? AND Password = ?", phone, pass).First(&employee).Error
	return employee, err
}
