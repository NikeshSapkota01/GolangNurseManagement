package controllers

import (
	"net/http"
	"strconv"

	"github.com/NikeshSapkota01/golangNurseManagement/config"
	"github.com/NikeshSapkota01/golangNurseManagement/core"
	"github.com/NikeshSapkota01/golangNurseManagement/models"
	"github.com/NikeshSapkota01/golangNurseManagement/services"
	"github.com/gin-gonic/gin"
)

type User struct {
	core.Controller
}

func (a *User) Index(c *gin.Context) {
	var params *models.User
	c.BindJSON(&params)
	// c.ShouldBind(&params)
	result, list := new(services.Users).All(params)
	a.Json(c, result, list)
}

func (a *User) Store(c *gin.Context) {
	var params *models.User
	c.BindJSON(&params)
	result, data := new(services.Users).Create(params)
	a.Json(c, result, data)
}

func (a *User) Update(c *gin.Context) {
	var params *models.User
	c.BindJSON(&params)
	result, data := new(services.Users).Update(params)
	a.Json(c, result, data)
}

func (a *User) Destroy(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err == nil {
		result, data := new(services.Users).Delete(id)
		a.Json(c, result, data)
		return
	}
	a.JsonFail(c, http.StatusBadRequest, "Not have params")
}

func (a *User) Show(c *gin.Context) {
	var account models.User
	result := config.Db().First(&account, c.Params.ByName("id"))
	a.Json(c, result, account)
}
