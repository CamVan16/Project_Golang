package controllers

import (
	"camvan/models"
	"camvan/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DepartmentController struct {
	service        services.DepartmentService
	subDeptService services.SubDepartmentService
}

func NewDepartmentController(service services.DepartmentService, subDeptService services.SubDepartmentService) *DepartmentController {
	return &DepartmentController{service, subDeptService}
}

func (ctrl *DepartmentController) CreateDepartment(c *gin.Context) {
	var department models.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.CreateDepartment(&department); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, department)
}

func (ctrl *DepartmentController) GetDepartmentByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	department, err := ctrl.service.GetDepartmentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, department)
}

func (ctrl *DepartmentController) GetAllDepartments(c *gin.Context) {
	departments, err := ctrl.service.GetAllDepartments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, departments)
}

func (ctrl *DepartmentController) UpdateDepartment(c *gin.Context) {
	var department models.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	department.ID = uint(id)
	if err := ctrl.service.UpdateDepartment(&department); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	subDepartments, err := ctrl.subDeptService.GetSubDepartmentsByDepartmentID(department.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, subDept := range subDepartments {
		subDept.DepartmentID = department.ID
		if err := ctrl.subDeptService.UpdateSubDepartment(&subDept); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, department)
}

func (ctrl *DepartmentController) DeleteDepartment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	subDepartments, err := ctrl.subDeptService.GetSubDepartmentsByDepartmentID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, subDept := range subDepartments {
		if err := ctrl.subDeptService.DeleteSubDepartment(subDept.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if err := ctrl.service.DeleteDepartment(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, true)
}
