package controller

import (
	"github.com/gin-gonic/gin"
	"go/note/util"
)

func CaptchaHandle(c *gin.Context) {
	id, captcha, err := util.GetCaptcha()
	if err != nil {
		c.JSON(500, gin.H{"code": 400, "errors": map[string]any{"body": []string{"服务器异常，请重试"}}})
	}
	c.JSON(200, gin.H{"code": 200, "msg": "验证码获取成功", "captcha": map[string]any{"captchaId": id, "captchaImg": captcha}})
}
