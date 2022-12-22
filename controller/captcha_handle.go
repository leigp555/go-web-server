package controller

import (
	"github.com/gin-gonic/gin"
	"go/note/util"
)

func CaptchaHandle(c *gin.Context) {
	id, captcha, err := util.GetCaptcha()
	if err != nil {
		c.JSON(500, gin.H{"msg": "服务器异常，请重试"})
	}
	c.JSON(200, gin.H{"captchaId": id, "captcha": captcha})
}
