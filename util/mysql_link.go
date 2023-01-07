package util

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"img.server/config"
	"log"
)

type myDb struct {
	Db *gorm.DB
}

var Mydb = new(myDb)

// LinkMysqlDB LinkDB 连接mysql数据库
func (db *myDb) LinkMysqlDB() *gorm.DB {
	dsn := fmt.Sprintf("root:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True", config.GlobalConfig.Mysql.Password, config.GlobalConfig.Mysql.Addr, config.GlobalConfig.Mysql.DB)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //配置一个日志
	})
	if err != nil {
		log.Panicf("mysql数据库连接失败%v\n", err)
		return nil
	}
	sqlDb, _ := d.DB()
	//设置连接池
	sqlDb.SetMaxIdleConns(config.GlobalConfig.Mysql.MaxConn) //设置最大连接数
	sqlDb.SetMaxOpenConns(config.GlobalConfig.Mysql.MaxConn) //设置最大的空闲连接数
	Mydb.Db = d
	fmt.Println("成功连接mysql数据库")
	return d
}
