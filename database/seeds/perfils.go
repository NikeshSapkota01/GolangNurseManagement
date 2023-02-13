package seeds

import (
	"log"

	"github.com/NikeshSapkota01/golangNurseManagement/config"
	"github.com/NikeshSapkota01/golangNurseManagement/models"
)

func Perfils() {
	db := config.Db()
	var Perfils []models.Perfil

	result := db.Find(&Perfils)
	if result.RowsAffected >= 4 {
		log.Printf("Seed (Perfils): Nothing to seed")
		return
	}

	Perfils = []models.Perfil{
		{Description: "ADMIN"},
		{Description: "GUEST"},
	}

	db.Create(&Perfils)
	log.Printf("Seed (Perfils): Success")
}
