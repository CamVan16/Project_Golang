package routes

import (
	"camvan/connection"
	"camvan/controllers"
	"camvan/repositories"
	"camvan/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	db := connection.DB

	departmentRepository := repositories.NewDepartmentRepository(db)
	departmentService := services.NewDepartmentService(departmentRepository)

	subDepartmentRepository := repositories.NewSubDepartmentRepository(db)
	subDepartmentService := services.NewSubDepartmentService(subDepartmentRepository)
	subDepartmentController := controllers.NewSubDepartmentController(subDepartmentService)
	departmentController := controllers.NewDepartmentController(departmentService, subDepartmentService)

	employeeRepository := repositories.NewEmployeeRepository(db)
	employeeService := services.NewEmployeeService(employeeRepository)
	employeeController := controllers.NewEmployeeController(employeeService)

	router.POST("/api/departments", departmentController.CreateDepartment)
	router.GET("/api/departments", departmentController.GetAllDepartments)
	router.GET("/api/departments/:id", departmentController.GetDepartmentByID)
	router.PUT("/api/departments/:id", departmentController.UpdateDepartment)
	router.DELETE("/api/departments/:id", departmentController.DeleteDepartment)

	router.POST("/api/sub_departments", subDepartmentController.CreateSubDepartment)
	router.GET("/api/sub_departments", subDepartmentController.GetAllSubDepartments)
	router.GET("/api/sub_departments/:id", subDepartmentController.GetSubDepartmentByID)
	router.PUT("/api/sub_departments/:id", subDepartmentController.UpdateSubDepartment)
	router.DELETE("/api/sub_departments/:id", subDepartmentController.DeleteSubDepartment)

	router.POST("/api/employees", employeeController.CreateEmployee)
	router.GET("/api/employees", employeeController.GetAllEmployees)
	router.GET("/api/employees/:id", employeeController.GetEmployeeByID)
	router.PUT("/api/employees/:id", employeeController.UpdateEmployee)
	router.DELETE("/api/employees/:id", employeeController.DeleteEmployee)

	return router
}
