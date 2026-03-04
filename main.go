package main

import (
	"context"
	"fmt"
	"hands-on-redis/config"
	"os"
)

func main() {

	fmt.Println(os.Getwd())

	err := config.InitConfigFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	if err := config.Init(); err != nil {
		panic(err)
	}

	rediStatus := config.RedisClient.Ping(context.Background())
	if rediStatus.Err() != nil {
		panic(rediStatus.Err())
	}

	fmt.Println(rediStatus.String())

}
