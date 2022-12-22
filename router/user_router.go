package router

import (
	"github.com/gin-gonic/gin"
	"go/note/controller"
)

func UserRouter(r *gin.RouterGroup) {
	r.GET("/login", controller.LoginHandle)
	r.GET("/register", controller.RegisterHandle)
}
