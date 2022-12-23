package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go/note/config"
	"go/note/middleware"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func InitRouter() {
	//gin配置log文件
	f, _ := os.OpenFile("log/log", os.O_RDWR|os.O_APPEND, 0755)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	r.Use(middleware.Cors())
	g := r.Group("v1/api")
	//注册路由
	UserRouter(g)
	ArticleRouter(g)
	CaptchaRouter(g)
	EmailCaptchaRouter(g)
	//监听端口
	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", config.GlobalConfig.Port),
		Handler: r,
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
