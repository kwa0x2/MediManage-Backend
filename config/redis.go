package config

import (
	"github.com/gin-contrib/sessions"
	redisStore "github.com/gin-contrib/sessions/redis"
	"github.com/go-redis/redis/v8"
	"os"
)

func RedisSession() redisStore.Store {
	store, err := redisStore.NewStore(10, "tcp", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASSWORD"), []byte(os.Getenv("SESSION_SECRET_KEY")))
	if err != nil {
		panic(err)
	}

	store.Options(sessions.Options{
		MaxAge:   2700,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	})

	return store
}

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	return rdb
}
