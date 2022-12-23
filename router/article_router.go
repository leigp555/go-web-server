package router

import (
	"github.com/gin-gonic/gin"
	"go/note/middleware"
)

func ArticleRouter(router *gin.RouterGroup) {
	//需要权限的路由组,添加token中间件验证
	r := router.Group("").Use(middleware.TokenVerify())
	{
		r.GET("/article", func(c *gin.Context) {
			userId, exists := c.Get("userId")
			if !exists {
				c.JSON(500, gin.H{"msg": "服务器异常请重试"})
			}
			c.JSON(200, userId)
		})

		r.GET("/xxx", func(c *gin.Context) {
			c.String(200, "register")
		})
	}

}
