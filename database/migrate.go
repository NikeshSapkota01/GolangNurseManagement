package database

import (
	"log"

	"github.com/NikeshSapkota01/golangNurseManagement/config"
	"github.com/NikeshSapkota01/golangNurseManagement/models"
)

func Migrate() {
	log.Printf("Migrate: Start")
	db := config.Db()
	db.AutoMigrate(
		&models.User{},
	)
	log.Printf("Migrate: Success")
}
