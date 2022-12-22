package router

import (
	"github.com/gin-gonic/gin"
	"go/note/controller"
)

func UserRouter(router *gin.RouterGroup) {
	r := router.Group("")
	{
		r.GET("/login", controller.LoginHandle)
		r.GET("/register", controller.RegisterHandle)
	}

}
