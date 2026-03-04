package config

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var (
	Params      = viper.New()
	RedisClient *redis.Client
)

func Init() error {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", Params.GetString("redis.address"), Params.GetString("redis.port")),
		DB:           viper.GetInt("redis.db"),
		DialTimeout:  time.Second * 3,
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 5,
	})
	fmt.Println("connect to redis")

	return nil
}

func InitConfigFile(configPath string) error {
	Params.SetConfigFile(configPath)
	viper.SetConfigName("config")
	viper.SetConfigName("yaml")
	if err := Params.ReadInConfig(); err != nil {
		return err
	}
	fmt.Println("config file loaded")
	return nil
}
