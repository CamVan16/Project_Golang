package controllers

import (
	"camvan/models"
	"camvan/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	service services.EmployeeService
}

func NewEmployeeController(service services.EmployeeService) *EmployeeController {
	return &EmployeeController{service}
}

func (ctrl *EmployeeController) CreateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.CreateEmployee(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (ctrl *EmployeeController) GetEmployeeByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	employee, err := ctrl.service.GetEmployeeByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (ctrl *EmployeeController) GetAllEmployees(c *gin.Context) {
	employees, err := ctrl.service.GetAllEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employees)
}

func (ctrl *EmployeeController) UpdateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	employee.ID = uint(id)
	if err := ctrl.service.UpdateEmployee(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (ctrl *EmployeeController) DeleteEmployee(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.service.DeleteEmployee(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, true)
}
