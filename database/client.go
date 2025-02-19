package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var ctx = context.Background()
var RedisClient *redis.Client

func Connect(addr, password string, db int) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Cannot connect to Redis: ", err)
	}
	log.Println("Connected to Redis!")
}
