package redis

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/garyburd/redigo/redis"
)

func TestInitRedisPool(t *testing.T) {
	host := "127.0.0.1"
	port := 6379
	pool := InitRedisPool(host, port)

	go func(*redis.Pool) {
		conn := pool.Get()
		defer conn.Close()
		for i := 0; i < 10; i++ {
			time.Sleep(2 * time.Second)
			_, err := conn.Do("rpush", "mqtest", "value"+strconv.Itoa(i))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}(pool)

	go func(*redis.Pool) {
		conn := pool.Get()
		defer conn.Close()
		for {
			v, err := redis.String(conn.Do("lpop", "mqtest"))
			if err != nil {
				fmt.Println("no val")
				time.Sleep(time.Second)
			} else {
				fmt.Printf("-- %s\n", v)
			}
		}
	}(pool)

	time.Sleep(10 * time.Second)
}
