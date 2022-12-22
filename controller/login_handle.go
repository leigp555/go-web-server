package controller

import (
	"github.com/gin-gonic/gin"
	"go/note/model"
	"go/note/util"
)

func LoginHandle(c *gin.Context) {
	var db = util.Mydb.Db
	user := model.User{
		Username: "lgp",
		Password: "123456",
		Email:    "907090585@qq.com",
	}
	db.Find(&user)
	c.JSON(200, user)
}
