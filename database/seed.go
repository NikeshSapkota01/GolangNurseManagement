package database

import (
	"log"

	"github.com/NikeshSapkota01/golangNurseManagement/database/seeds"
)

func Seeder() {
	log.Printf("Seed: Start")
	seeds.Perfils()
	log.Printf("Seed: Success")
}
