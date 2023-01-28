package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"img.server/model"
	"img.server/util"
)

func RegisterHandle(c *gin.Context) {
	type reqData struct {
		Username         string `json:"username" binding:"required,min=1,max=20" msg:"用户名不能为空,且长度为1~20位"`
		Email            string `json:"email" binding:"required,email" msg:"请输入正确的邮箱"`
		Password         string `json:"password" binding:"required,min=6,max=12" msg:"密码不能为空,且长度为6~12位"`
		CheckPassword    string `json:"checkPassword" binding:"required,min=6,max=12,eqfield=Password" msg:"两次输入的密码不一致"`
		EmailCaptchaCode string `json:"emailCaptcha" binding:"required,len=6" msg:"邮箱验证码不正确"`
		CaptchaId        string `json:"captchaId" binding:"required" msg:"图形验证码不正确"`
		Captcha          string `json:"captcha" binding:"required" msg:"图形验证码不正确"`
	}
	//json数据验证不通过
	var userInfo reqData
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		msg := util.GetValidMsg(err, &userInfo)
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": msg}})
		return
	}
	//邮箱验证码验证
	var rdb = util.Redb.Db
	var ctx = context.Background()
	val, err2 := rdb.Get(ctx, userInfo.Email).Result()
	if err2 != nil || val != userInfo.EmailCaptchaCode {
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": "邮箱验证码不正确"}})
		return
	}
	//图形验证码验证
	isRight := util.VerifyCaptcha(userInfo.CaptchaId, userInfo.Captcha)
	if !isRight {
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": "图形验证码错误"}})
		return
	}

	//验证数据库中是否存在该用户
	var mdb = util.Mydb.Db
	var u = model.User{}
	mdb.Where("email = ?", userInfo.Email).Find(&u)
	if u.Email == userInfo.Email {
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": "用户已存在"}})
		return
	}

	//数据验证通过,将用户信息保存在数据库中
	user := model.User{
		Username: userInfo.Username,
		Password: util.Md5Str(userInfo.Password),
		Email:    userInfo.Email,
	}
	mdb.Create(&user)
	c.JSON(200, gin.H{"code": 200, "msg": "注册成功", "user": map[string]any{"email": userInfo.Email, "username": userInfo.Username}})
}
