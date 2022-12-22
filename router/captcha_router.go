package router

import (
	"github.com/gin-gonic/gin"
	"go/note/controller"
)

func CaptchaRouter(r *gin.RouterGroup) {
	r.GET("/captcha", controller.CaptchaHandle)
}
