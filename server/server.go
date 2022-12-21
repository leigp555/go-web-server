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
	f, _ := os.OpenFile("log/log", os.O_RDWR|os.O_APPEND, 0755)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", config.GlobalConfig.Port),
		Handler: router,
	}

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
