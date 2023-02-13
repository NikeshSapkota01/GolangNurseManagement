package routes

import (
	"github.com/NikeshSapkota01/golangNurseManagement/controllers"
	"github.com/gin-gonic/gin"
)

func V1(r *gin.Engine) *gin.Engine {

	v1 := r.Group("/api/v1")
	{
		user := new(controllers.User)
		v1.GET("/users", user.Index)
		v1.POST("/users", user.Store)
		v1.PATCH("/users", user.Update)
		v1.DELETE("/users/:id", user.Destroy)
		v1.GET("/users/:id", user.Show)
	}

	return r
}
