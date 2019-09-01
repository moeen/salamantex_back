package redis

import (
	"gihub.com/moeen/salamantex_back/config"
	"github.com/go-redis/redis"
	"log"
)

var redisClient *redis.Client

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: config.GetConfig().Redis.Address,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatalln(err)
	}
}

func GetRedis() *redis.Client {
	return redisClient
}
