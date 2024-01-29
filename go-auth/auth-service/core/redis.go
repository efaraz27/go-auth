package core

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func ConnetRedis(host string, port int, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       0, // use default DB
	})

	return client
}
