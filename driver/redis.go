package driver

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gopkg.in/ini.v1"
	"log"
)

var RDB *redis.Client
var ctx = context.Background()

func init() {
	config,err := ini.Load("./config/config.ini")
	if err != nil {
		log.Fatal(err)
		return
	}

	ip := config.Section("redis").Key("ip").String()
	port := config.Section("redis").Key("port").String()

	RDB = redis.NewClient(&redis.Options{
		Addr: ip + ":" + port,
		Password: "",
		DB: 10,
	})

	_,err = RDB.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

}