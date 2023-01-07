package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"img.server/util"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func EmailCaptchaHandle(c *gin.Context) {
	type UserEmail struct {
		Email string `json:"email" binding:"required,email" msg:"邮箱格式不正确"`
	}
	var userEmail UserEmail

	//json验证
	err := c.ShouldBindJSON(&userEmail)
	if err != nil {
		msg := util.GetValidMsg(err, &userEmail)
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": []string{msg}}})
		return
	}
	//验证码发送
	var arr = make([]string, 0)
	for i := 0; i < 6; i++ {
		arr = append(arr, strconv.Itoa(rand.Intn(10)))
	}
	randStr := fmt.Sprintf(strings.Join(arr, ""))
	err = util.SendEmail([]string{userEmail.Email}, randStr)
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "errors": map[string]any{"body": []string{"验证码发送失败,请重试"}}})
		return
	}
	//存入redis
	var rdb = util.Redb.Db
	var ctx = context.Background()
	err = rdb.Set(ctx, userEmail.Email, randStr, 300*time.Second).Err()
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "errors": map[string]any{"body": []string{"服务器异常请重试"}}})
		return
	}
	//返回成功的响应
	c.JSON(200, gin.H{"code": 200, "msg": "验证码发送成功,请前往邮箱查看", "captcha": map[string]any{}})
}
