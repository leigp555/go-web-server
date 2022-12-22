package controller

import "github.com/gin-gonic/gin"

func RegisterHandle(c *gin.Context) {
	c.String(200, "register")
}
