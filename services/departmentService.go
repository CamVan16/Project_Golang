package services

import (
	"camvan/models"
	"camvan/repositories"
)

type DepartmentService interface {
	CreateDepartment(department *models.Department) error
	GetDepartmentByID(id uint) (models.Department, error)
	GetAllDepartments() ([]models.Department, error)
	UpdateDepartment(department *models.Department) error
	DeleteDepartment(id uint) error
}

type departmentService struct {
	repository repositories.DepartmentRepository
}

func NewDepartmentService(repository repositories.DepartmentRepository) DepartmentService {
	return &departmentService{repository}
}

func (s *departmentService) CreateDepartment(department *models.Department) error {
	return s.repository.Create(department)
}

func (s *departmentService) GetDepartmentByID(id uint) (models.Department, error) {
	return s.repository.FindByID(id)
}

func (s *departmentService) GetAllDepartments() ([]models.Department, error) {
	return s.repository.FindAll()
}

func (s *departmentService) UpdateDepartment(department *models.Department) error {
	return s.repository.Update(department)
}

func (s *departmentService) DeleteDepartment(id uint) error {
	return s.repository.Delete(id)
}
