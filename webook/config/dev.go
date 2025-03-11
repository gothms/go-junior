//go:build !k8s

package config

var Config = config{
	DB: DBConfig{
		DSN: "root:root@tcp(localhost:13316)/webook", // docker
		//DSN: "root:1234@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local", // 本地
	},
	Redis: RedisConfig{
		Addr: "localhost:6379",
	},
}
