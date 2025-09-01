package config

import (
	"github.com/redis/go-redis/v9"
	"context"
	"github.com/joho/godotenv"
	"os"


)

var (
	RedisClient *redis.Client
	Ctx = context.Background()
)

func InitRedis() {
	godotenv.Load()
	RedisClient = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: os.Getenv("REDIS_PASSWORD"),
			DB: 0,
		})
}