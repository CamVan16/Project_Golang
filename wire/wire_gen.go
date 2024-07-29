// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"camvan/controllers"
	"camvan/repositories"
	"camvan/services"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeDepartment(db *gorm.DB, subDeptService services.SubDepartmentService) *controllers.DepartmentController {
	departmentRepository := repositories.NewDepartmentRepository(db)
	departmentService := services.NewDepartmentService(departmentRepository)
	departmentController := controllers.NewDepartmentController(departmentService, subDeptService)
	return departmentController
}

func InitializeSubDepartment(db *gorm.DB) *controllers.SubDepartmentController {
	subDepartmentRepository := repositories.NewSubDepartmentRepository(db)
	subDepartmentService := services.NewSubDepartmentService(subDepartmentRepository)
	subDepartmentController := controllers.NewSubDepartmentController(subDepartmentService)
	return subDepartmentController
}

func InitializeEmployee(db *gorm.DB) *controllers.EmployeeController {
	employeeRepository := repositories.NewEmployeeRepository(db)
	employeeService := services.NewEmployeeService(employeeRepository)
	employeeController := controllers.NewEmployeeController(employeeService)
	return employeeController
}

func InitializeUser(db *gorm.DB, employeeRepo repositories.EmployeeRepository) *controllers.UserController {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository, employeeRepo)
	userController := controllers.NewUserController(userService)
	return userController
}
