package controller

import "github.com/gin-gonic/gin"

func LoginHandle(c *gin.Context) {
	c.String(200, "login")
}
