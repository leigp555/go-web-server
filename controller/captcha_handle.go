package controller

import (
	"github.com/gin-gonic/gin"
	"img.server/util"
)

func CaptchaHandle(c *gin.Context) {
	id, captcha, err := util.GetCaptcha()
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "errors": map[string]any{"body": []string{"服务器异常，请重试"}}})
	}
	c.JSON(200, gin.H{"code": 200, "msg": "验证码获取成功", "data": map[string]any{"captchaId": id, "captchaImg": captcha}})
}
