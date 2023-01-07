package router

import (
	"github.com/gin-gonic/gin"
	"img.server/controller"
)

func EmailCaptchaRouter(r *gin.RouterGroup) {
	r.POST("/email/captcha", controller.EmailCaptchaHandle)
}
