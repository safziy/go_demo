package redis_demo

import (
	"github.com/go-redis/redis/v8"
	"log"
)

var Key1 string = "Book:Name"
var Value1 string = "The Go Programming Language"

func RedisStringDemo() {
	client := GetRedisConnection()
	defer func(client *redis.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalf("Close redis connection error {%v}", err)
		}
	}(client)
	RedisStringSet(client)
	RedisStringGet(client)
}

func RedisStringSet(client *redis.Client) {
	_, err := client.Set(Ctx, Key1, Value1, 0).Result()
	if err != nil {
		log.Fatalf("Redis command {SET} error {%v}", err)
	}
}

func RedisStringGet(client *redis.Client) {
	_, err := client.Get(Ctx, Key1).Result()
	if err != nil {
		log.Fatalf("Close command {GET} error {%v}", err)
	}
}
