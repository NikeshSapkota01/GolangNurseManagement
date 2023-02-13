package models

import (
	"github.com/NikeshSapkota01/golangNurseManagement/core"
)

type Perfil struct {
	Name        string `json:"name" gorm:"index;not null"`
	Description string `json:"description" gorm:"not null"`
	core.Model
}
