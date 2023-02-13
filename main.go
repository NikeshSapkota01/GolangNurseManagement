package main

import (
	"github.com/NikeshSapkota01/golangNurseManagement/database"
	"github.com/NikeshSapkota01/golangNurseManagement/routes"
)

func main() {
	database.Migrate()
	// database.Seeder()
	routes.Run()
}
