package redis_demo

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
)

var ctx = context.Background()

func TestRedisConnection()  {
	client := redis.NewClient(&redis.Options{
		Addr: "1127.0.0.1:6379",
		Password: "",
		DB: 0,
	})

	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)
}
