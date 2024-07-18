package services

import (
	"camvan/models"
	"camvan/repositories"
)

type EmployeeService interface {
	CreateEmployee(employee *models.Employee) error
	GetEmployeeByID(id uint) (models.Employee, error)
	GetAllEmployees() ([]models.Employee, error)
	UpdateEmployee(employee *models.Employee) error
	DeleteEmployee(id uint) error
}

type employeeService struct {
	repository repositories.EmployeeRepository
}

func NewEmployeeService(repository repositories.EmployeeRepository) EmployeeService {
	return &employeeService{repository}
}

func (s *employeeService) CreateEmployee(employee *models.Employee) error {
	return s.repository.Create(employee)
}

func (s *employeeService) GetEmployeeByID(id uint) (models.Employee, error) {
	return s.repository.FindByID(id)
}

func (s *employeeService) GetAllEmployees() ([]models.Employee, error) {
	return s.repository.FindAll()
}

func (s *employeeService) UpdateEmployee(employee *models.Employee) error {
	return s.repository.Update(employee)
}

func (s *employeeService) DeleteEmployee(id uint) error {
	return s.repository.Delete(id)
}
