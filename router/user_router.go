package router

import (
	"github.com/gin-gonic/gin"
	"img.server/controller"
)

func UserRouter(router *gin.RouterGroup) {
	r := router.Group("user")
	{
		r.POST("/login", controller.LoginHandle)
		r.POST("/register", controller.RegisterHandle)
	}

}
