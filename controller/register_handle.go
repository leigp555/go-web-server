package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"go/note/model"
	"go/note/util"
)

func RegisterHandle(c *gin.Context) {
	type UserInfo struct {
		Username         string `json:"username" binding:"required,min=3,max=20" msg:"用户名不能为空,且长度为3~20位"`
		Password         string `json:"password" binding:"required,min=6,max=12" msg:"密码不能为空,且长度为6~12位"`
		RePassword       string `json:"re_password" binding:"required,min=6,max=12,eqfield=Password" msg:"两次输入的密码不一致"`
		Email            string `json:"email" binding:"required,email" msg:"请输入正确的邮箱"`
		EmailCaptchaCode string `json:"email_captcha_code" binding:"required,len=6" msg:"邮箱验证码不正确"`
		CaptchaId        string `json:"captcha_id" binding:"required" msg:"图形验证码不正确"`
		CaptchaCode      string `json:"captcha_code" binding:"required" msg:"图形验证码不正确"`
	}
	//json数据验证不通过
	var userInfo UserInfo
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		msg := util.GetValidMsg(err, &userInfo)
		c.JSON(400, gin.H{"code": 400, "errMsg": msg})
		return
	}
	//邮箱验证码验证
	var rdb = util.Redb.Db
	var ctx = context.Background()
	val, err2 := rdb.Get(ctx, userInfo.Email).Result()
	if err2 != nil || val != userInfo.EmailCaptchaCode {
		c.JSON(400, gin.H{"code": 400, "errMsg": "邮箱验证码不正确"})
		return
	}
	//图形验证码验证
	isRight := util.VerifyCaptcha(userInfo.CaptchaId, userInfo.CaptchaCode)
	if !isRight {
		c.JSON(400, gin.H{"code": 400, "errMsg": "图形验证码错误"})
		return
	}
	//数据验证通过,将用户信息保存在数据库中
	user := model.User{
		Username: userInfo.Username,
		Password: userInfo.Password,
		Email:    userInfo.Email,
	}
	var mdb = util.Mydb.Db
	mdb.Create(&user)
	c.JSON(200, gin.H{"code": 200, "msg": "注册成功"})
}
