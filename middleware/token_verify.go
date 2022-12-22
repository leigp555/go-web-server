package middleware

import (
	"github.com/gin-gonic/gin"
	"go/note/model"
	"go/note/util"
	"strings"
)

func TokenVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		var mdb = util.Mydb.Db
		//从请求头获取token
		tokenHeader := c.GetHeader("Authorization")

		//从请求头获取token失败
		if tokenHeader == "" {
			c.JSON(401, gin.H{"msg": "请上传身份凭证"})
			return
		}
		//拆分出token
		splitArr := strings.Split(tokenHeader, " ")
		tokenStr := splitArr[1]
		//解析token 解析失败阻止后续中间件执行
		userId, err2 := util.ParseToken(tokenStr)
		if err2 != nil {
			c.JSON(403, gin.H{"msg": "用户身份过期,请重新登录"})
			c.Abort()
			return
		}
		var dbUser = model.User{}
		mdb.Where("id = ?", userId).First(&dbUser)
		c.Set("userId", dbUser.ID)
		c.Set("userEmail", dbUser.Email)
	}
}
