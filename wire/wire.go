package wire

import (
	"camvan/controllers"
	"camvan/repositories"
	"camvan/services"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeDepartment1(db *gorm.DB, subDeptService services.SubDepartmentService) *controllers.DepartmentController {
	wire.Build(repositories.NewDepartmentRepository,
		services.NewDepartmentService,
		controllers.NewDepartmentController)
	return &controllers.DepartmentController{}
}

func InitializeSubDepartment1(db *gorm.DB) *controllers.SubDepartmentController {
	wire.Build(repositories.NewSubDepartmentRepository,
		services.NewSubDepartmentService,
		controllers.NewSubDepartmentController)
	return &controllers.SubDepartmentController{}
}

func InitializeEmployee1(db *gorm.DB) *controllers.EmployeeController {
	wire.Build(repositories.NewEmployeeRepository,
		services.NewEmployeeService,
		controllers.NewEmployeeController)
	return &controllers.EmployeeController{}
}

func InitializeUser1(db *gorm.DB, employeeRepo repositories.EmployeeRepository) *controllers.UserController {
	wire.Build(repositories.NewUserRepository,
		services.NewUserService,
		controllers.NewUserController)
	return &controllers.UserController{}
}
