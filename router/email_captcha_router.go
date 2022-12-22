package router

import (
	"github.com/gin-gonic/gin"
	"go/note/controller"
)

func EmailCaptchaRouter(r *gin.RouterGroup) {
	r.GET("/email/captcha", controller.EmailCaptchaHandle)
}
