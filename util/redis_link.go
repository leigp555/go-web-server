package util

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"img.server/config"
	"log"
	"time"
)

type reDb struct {
	Db *redis.Client
}

var Redb = new(reDb)

// LinkRedisDB连接redis数据库

func (*reDb) LinkRedisDB() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.GlobalConfig.Redis.Addr,
		Password: config.GlobalConfig.Redis.Password,
		DB:       config.GlobalConfig.Redis.DB,       // use default DB
		PoolSize: config.GlobalConfig.Redis.PoolSize, // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Panicf("redis数据库连接失败%v\n", err)
		return nil
	}
	Redb.Db = rdb
	fmt.Println("成功连接redis数据库")
	return rdb
}
