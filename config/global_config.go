package config

import (
	"time"
)

type mysqlConfig struct {
	Addr     string
	Password string
	MaxConn  int
	MaxOpen  int
	DB       string
}

type redisConfig struct {
	Addr     string
	Password string
	DB       int
	PoolSize int
}

type tokenConfig struct {
	SigningKey  string
	ExpiresTime time.Time
}

type globalConfig struct {
	Mysql mysqlConfig
	Redis redisConfig
	Port  string
	Token tokenConfig
}

var GlobalConfig = globalConfig{
	Mysql: mysqlConfig{
		Addr:     "1.117.141.66:3306",
		Password: "lgp1234567",
		MaxConn:  200,
		MaxOpen:  100,
		DB:       "img",
	},
	Redis: redisConfig{
		Addr:     "1.117.141.66:6379",
		Password: "123456abc",
		DB:       0,
		PoolSize: 50,
	},
	Port: "8080",
	Token: tokenConfig{
		SigningKey:  "e199ad17-c090-43cb-b095-1dc55c209a77",
		ExpiresTime: time.Now().Add(1000 * time.Second),
	},
}
