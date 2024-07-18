package services

import (
	"camvan/models"
	"camvan/repositories"
)

type SubDepartmentService interface {
	CreateSubDepartment(subDepartment *models.SubDepartment) error
	GetSubDepartmentByID(id uint) (models.SubDepartment, error)
	GetAllSubDepartments() ([]models.SubDepartment, error)
	GetSubDepartmentsByDepartmentID(departmentID uint) ([]models.SubDepartment, error)
	UpdateSubDepartment(subDepartment *models.SubDepartment) error
	DeleteSubDepartment(id uint) error
}

type subDepartmentService struct {
	repository repositories.SubDepartmentRepository
}

func NewSubDepartmentService(repository repositories.SubDepartmentRepository) SubDepartmentService {
	return &subDepartmentService{repository}
}

func (s *subDepartmentService) CreateSubDepartment(subDepartment *models.SubDepartment) error {
	return s.repository.Create(subDepartment)
}

func (s *subDepartmentService) GetSubDepartmentByID(id uint) (models.SubDepartment, error) {
	return s.repository.FindByID(id)
}

func (s *subDepartmentService) GetAllSubDepartments() ([]models.SubDepartment, error) {
	return s.repository.FindAll()
}

func (s *subDepartmentService) GetSubDepartmentsByDepartmentID(departmentID uint) ([]models.SubDepartment, error) {
	return s.repository.FindByDepartmentID(departmentID)
}

func (s *subDepartmentService) UpdateSubDepartment(subDepartment *models.SubDepartment) error {
	return s.repository.Update(subDepartment)
}

func (s *subDepartmentService) DeleteSubDepartment(id uint) error {
	return s.repository.Delete(id)
}
