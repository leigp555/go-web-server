package config

type mysqlConfig struct {
	Addr     string
	Password string
}

type redisConfig struct {
	Addr     string
	Password string
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
	},
	Redis: redisConfig{
		Addr:     "1.117.141.66:6379",
		Password: "123456abc",
	},
	Port: "8080",
}
