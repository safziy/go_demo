package redis_demo

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var Ctx = context.Background()

func GetRedisConnection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Redis get connection error {%v}", err)
	}
	return client
}

func RedisDemo() {
	RedisStringDemo()
	log.Println("All Redis Command done!")
}
