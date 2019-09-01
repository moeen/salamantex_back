package main

import (
	"gihub.com/moeen/salamantex_back/config"
	"gihub.com/moeen/salamantex_back/models"
	"gihub.com/moeen/salamantex_back/processor"
	r "gihub.com/moeen/salamantex_back/redis"
	"gihub.com/moeen/salamantex_back/server"
	"github.com/go-redis/redis"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"time"
)

func init() {
	config.SetConfig()
	r.InitRedis()
	models.InitDB()
}

func main() {
	defer models.GetDB().Close()
	defer r.GetRedis().Close()

	go server.InitServer()

	for {
		tx, err := r.GetRedis().LPop(config.GetConfig().Redis.TxQueue).Result()

		if err != nil && err != redis.Nil {
			log.Println(err)
		} else if err == nil {
			go processor.ProcessTransaction(tx)
		}

		time.Sleep(30 * time.Second)
	}
}
