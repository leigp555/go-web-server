package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/note/model"
	"go/note/util"
	"strconv"
)

func LoginHandle(c *gin.Context) {
	type reqData struct {
		User struct {
			Email    string `json:"email" binding:"required,email" msg:"请输入正确的邮箱"`
			Password string `json:"password" binding:"required,min=6,max=12" msg:"密码不能为空,且长度为6~12位"`
		} `json:"user"`
		CaptchaId   string `json:"captcha_id" binding:"required" msg:"图形验证码不正确"`
		CaptchaCode string `json:"captcha_code" binding:"required" msg:"图形验证码不正确"`
	}
	//验证json数据绑定
	var userInfo reqData
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		msg := util.GetValidMsg(err, &userInfo)
		fmt.Println(msg)
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": []string{msg}}})
		return
	}
	//验证图形验证码
	isRight := util.VerifyCaptcha(userInfo.CaptchaId, userInfo.CaptchaCode)
	if !isRight {
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": []string{"图形验证码错误"}}})
		return
	}
	//数据库中查询是否存在该用户
	var mdb = util.Mydb.Db
	var u = model.User{}
	mdb.Where("email = ?", userInfo.User.Email).First(&u)
	if u.Email != userInfo.User.Email {
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": []string{"用户还未注册,请先注册"}}})
		return
	}
	//验证密码
	if u.Password != util.Md5Str(userInfo.User.Password) {
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": []string{"用户名或密码不正确"}}})
		return
	}
	//签发token
	token, err2 := util.GenerateToken(strconv.Itoa(int(u.ID)))
	if err2 != nil {
		c.JSON(500, gin.H{"code": 500, "errors": map[string]any{"body": []string{"服务器异常,请稍后再试"}}})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "登录成功", "token": token, "user": map[string]any{"email": u.Email, "username": u.Username}})
}
