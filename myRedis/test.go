package myRedis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func Connect() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	fmt.Println("redis conn success")
	_, err = c.Do("set", "key", 100)
	if err != nil {
		fmt.Println("set key error:", err)
	}
	//r, err := redis.Int(c.Do("get", "key"))
	//if err != nil {
	//	fmt.Println("get key error:", err)
	//}
	//fmt.Println("key:", r)
	//_, err = c.Do("del", "key")
	//if err != nil {
	//	fmt.Println("del key error:", err)
	//}
	_, err = c.Do("expire", "key", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
}
