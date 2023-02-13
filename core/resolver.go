package core

import (
	"github.com/NikeshSapkota01/golangNurseManagement/config"
	"gorm.io/gorm"
)

type Resolver struct {
}

func (r *Resolver) Db() *gorm.DB {
	return config.Db()
}
