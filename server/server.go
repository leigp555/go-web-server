package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go/note/config"
	"go/note/util"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func StartServer() {
	//连接mysql数据库
	_ = util.Mydb.LinkMysqlDB()
	// 连接redis数据库
	_ = util.Redb.LinkRedisDB()

	//gin配置log文件
	f, _ := os.OpenFile("log/log", os.O_RDWR|os.O_APPEND, 0755)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router := gin.Default()

	//注册路由
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})
	//获取token
	router.GET("/token", func(c *gin.Context) {
		token, err := util.GenerateToken("我是你的眼")
		if err != nil {
			c.JSON(400, gin.H{"msg": "token生成失败"})
		}
		c.String(http.StatusOK, token)
	})
	//解析token
	router.POST("/parse", func(c *gin.Context) {
		str, _ := c.GetQuery("token")
		username, err2 := util.ParseToken(str)
		if err2 != nil {
			c.JSON(400, gin.H{"msg": "token验证失败"})
		}
		c.String(http.StatusOK, username)
	})
	// 生成验证码
	router.GET("/captcha", func(c *gin.Context) {
		id, captcha, err := util.GetCaptcha()
		if err != nil {
			c.JSON(500, gin.H{"msg": "服务器异常，请重试"})
		}
		c.JSON(200, gin.H{"id": id, "captcha": captcha})
	})
	//解析验证码
	router.POST("/parseCaptcha", func(c *gin.Context) {
		id, _ := c.GetQuery("id")
		code, _ := c.GetQuery("code")
		ret := util.VerifyCaptcha(id, code)
		fmt.Println(ret)
		c.JSON(200, gin.H{"ret": "xxx"})
	})

	//监听端口
	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", config.GlobalConfig.Port),
		Handler: router,
	}

	fmt.Printf("成功监听%s端口", config.GlobalConfig.Port)
	//服务启停
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
