package config

type mysqlConfig struct {
	Addr     string
	Password string
	MaxConn  int
	MaxOpen  int
}

type redisConfig struct {
	Addr     string
	Password string
	DB       int
	PoolSize int
}

type globalConfig struct {
	Mysql mysqlConfig
	Redis redisConfig
	Port  string
}

var GlobalConfig = globalConfig{
	Mysql: mysqlConfig{
		Addr:     "127.0.0.1:3306",
		Password: "123456",
		MaxConn:  200,
		MaxOpen:  100,
	},
	Redis: redisConfig{
		Addr:     "1.117.141.66:6379",
		Password: "123456abc",
		DB:       0,
		PoolSize: 50,
	},
	Port: "8080",
}
