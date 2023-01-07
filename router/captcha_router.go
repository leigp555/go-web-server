package router

import (
	"github.com/gin-gonic/gin"
	"img.server/controller"
)

func CaptchaRouter(r *gin.RouterGroup) {
	r.GET("/captcha", controller.CaptchaHandle)
}
