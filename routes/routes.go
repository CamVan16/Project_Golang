package routes

import (
	"camvan/connection"
	"camvan/middleware"
	"camvan/repositories"
	"camvan/services"
	"camvan/wire"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	db := connection.DB

	var subDeptService services.SubDepartmentService
	var employeeRepo repositories.EmployeeRepository

	subDepartmentController := wire.InitializeSubDepartment(db)
	departmentController := wire.InitializeDepartment(db, subDeptService)
	employeeController := wire.InitializeEmployee(db)
	userController := wire.InitializeUser(db, employeeRepo)

	router.POST("/api/departments", departmentController.CreateDepartment)
	router.GET("/api/departments", departmentController.GetAllDepartments)
	router.GET("/api/departments/:id", departmentController.GetDepartmentByID)

	router.POST("/api/sub_departments", subDepartmentController.CreateSubDepartment)
	router.GET("/api/sub_departments", subDepartmentController.GetAllSubDepartments)
	router.GET("/api/sub_departments/:id", subDepartmentController.GetSubDepartmentByID)

	router.POST("/api/employees", employeeController.CreateEmployee)
	router.GET("/api/employees", employeeController.GetAllEmployees)
	router.GET("/api/employees/:id", employeeController.GetEmployeeByID)

	router.POST("/api/users", userController.SignUpUser)
	router.POST("/api/users/signin", userController.SignInUser)
	router.POST("/api/users/refresh", userController.RefreshToken)

	authRoutes := router.Group("/")
	authRoutes.Use(middleware.AuthMiddleware())
	authRoutes.PUT("/api/departments/:id", departmentController.UpdateDepartment)
	authRoutes.DELETE("/api/departments/:id", departmentController.DeleteDepartment)
	authRoutes.PUT("/api/sub_departments/:id", subDepartmentController.UpdateSubDepartment)
	authRoutes.DELETE("/api/sub_departments/:id", subDepartmentController.DeleteSubDepartment)
	authRoutes.PUT("/api/employees/:id", employeeController.UpdateEmployee)
	authRoutes.DELETE("/api/employees/:id", employeeController.DeleteEmployee)
	authRoutes.GET("/api/users", userController.GetAllUsers)
	authRoutes.DELETE("/api/user/:id", userController.DeleteUser)
	return router

	// employeeRepository := repositories.NewEmployeeRepository(db)
	// employeeService := services.NewEmployeeService(employeeRepository)
	// employeeController := controllers.NewEmployeeController(employeeService)
}
