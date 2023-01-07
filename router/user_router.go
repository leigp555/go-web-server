package router

import (
	"github.com/gin-gonic/gin"
	"img.server/controller"
)

func UserRouter(router *gin.RouterGroup) {
	r := router.Group("users")
	{
		r.GET("/login", controller.LoginHandle)
		r.GET("/", controller.RegisterHandle)
	}

}
