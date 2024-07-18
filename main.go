package main

import (
	"camvan/connection"
	"camvan/models"
	"camvan/routes"
)

func main() {
	connection.ConnectDatabase()
	connection.DB.AutoMigrate(&models.Department{}, &models.Employee{}, &models.SubDepartment{})
	router := routes.SetupRouter()
	router.Run("localhost:8080")
}
