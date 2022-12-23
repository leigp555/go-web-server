package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"go/note/model"
	"go/note/util"
)

func RegisterHandle(c *gin.Context) {
	type reqData struct {
		User struct {
			Email      string `json:"email" binding:"required,email" msg:"请输入正确的邮箱"`
			Password   string `json:"password" binding:"required,min=6,max=12" msg:"密码不能为空,且长度为6~12位"`
			RePassword string `json:"re_password" binding:"required,min=6,max=12,eqfield=Password" msg:"两次输入的密码不一致"`
		} `json:"user"`
		EmailCaptchaCode string `json:"email_captcha_code" binding:"required,len=6" msg:"邮箱验证码不正确"`
		CaptchaId        string `json:"captcha_id" binding:"required" msg:"图形验证码不正确"`
		CaptchaCode      string `json:"captcha_code" binding:"required" msg:"图形验证码不正确"`
	}
	//json数据验证不通过
	var userInfo reqData
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		msg := util.GetValidMsg(err, &userInfo)
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": []string{msg}}})
		return
	}
	//邮箱验证码验证
	var rdb = util.Redb.Db
	var ctx = context.Background()
	val, err2 := rdb.Get(ctx, userInfo.User.Email).Result()
	if err2 != nil || val != userInfo.EmailCaptchaCode {
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": []string{"邮箱验证码不正确"}}})
		return
	}
	//图形验证码验证
	isRight := util.VerifyCaptcha(userInfo.CaptchaId, userInfo.CaptchaCode)
	if !isRight {
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": []string{"图形验证码错误"}}})
		return
	}

	//验证数据库中是否存在该用户
	var mdb = util.Mydb.Db
	var u = model.User{}
	mdb.Where("email = ?", userInfo.User.Email).Find(&u)
	if u.Email == userInfo.User.Email {
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": []string{"用户已存在"}}})
		return
	}

	//数据验证通过,将用户信息保存在数据库中
	user := model.User{
		Password: userInfo.User.Password,
		Email:    userInfo.User.Email,
	}
	mdb.Create(&user)
	c.JSON(200, gin.H{"code": 200, "msg": "注册成功", "user": map[string]any{"email": userInfo.User.Email}})
}
