package server

import (
	"img.server/model"
	"img.server/router"
	"img.server/util"
)

func StartServer() {
	//连接mysql数据库
	_ = util.Mydb.LinkMysqlDB()
	// 连接redis数据库
	_ = util.Redb.LinkRedisDB()
	//创建mysql数据库表
	model.InitDb()
	//初始化路由,创建服务
	router.InitRouter()
}
