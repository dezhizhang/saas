package driver

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

var Conn redis.Conn

func init()  {
	conn,err := redis.Dial("tcp","0.0.0.0:6379")
	defer conn.Close()

	if err != nil{
		log.Fatal(err)
		return
	}
	Conn = conn
	fmt.Println("redis边接成功")

}