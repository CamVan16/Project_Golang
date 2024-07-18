package controllers

import (
	"camvan/models"
	"camvan/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubDepartmentController struct {
	service services.SubDepartmentService
}

func NewSubDepartmentController(service services.SubDepartmentService) *SubDepartmentController {
	return &SubDepartmentController{service}
}

func (ctrl *SubDepartmentController) CreateSubDepartment(c *gin.Context) {
	var subDepartment models.SubDepartment
	if err := c.ShouldBindJSON(&subDepartment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.CreateSubDepartment(&subDepartment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subDepartment)
}

func (ctrl *SubDepartmentController) GetSubDepartmentByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	subDepartment, err := ctrl.service.GetSubDepartmentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, subDepartment)
}

func (ctrl *SubDepartmentController) GetAllSubDepartments(c *gin.Context) {
	subDepartments, err := ctrl.service.GetAllSubDepartments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subDepartments)
}

func (ctrl *SubDepartmentController) UpdateSubDepartment(c *gin.Context) {
	var subDepartment models.SubDepartment
	if err := c.ShouldBindJSON(&subDepartment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	subDepartment.ID = uint(id)
	if err := ctrl.service.UpdateSubDepartment(&subDepartment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subDepartment)
}

func (ctrl *SubDepartmentController) DeleteSubDepartment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.service.DeleteSubDepartment(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, true)
}
