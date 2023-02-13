package models

import (
	"time"

	"github.com/NikeshSapkota01/golangNurseManagement/core"
)

type User struct {
	Name      string     `json:"name" gorm:"index;not null"`
	Email     string     `json:"email" gorm:"index;not null"`
	FirstName *string    `json:"firstName"`
	LastName  *string    `json:"lastName"`
	Date      *time.Time `json:"date"`
	Gender    *int       `json:"gender" binding:"required"`
	Password  string     `json:"password"`
	core.Model
}
