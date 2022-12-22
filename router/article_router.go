package router

import "github.com/gin-gonic/gin"

func ArticleRouter(r *gin.RouterGroup) {
	r.GET("/article", func(c *gin.Context) {
		c.String(200, "login")
	})
	r.GET("/xxx", func(c *gin.Context) {
		c.String(200, "register")
	})
}
