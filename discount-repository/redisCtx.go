package discountrepository

import (
	"context"
	"fmt"

	"github.com/alijkdkar/ArvanChallenge/pkg"
	"github.com/redis/go-redis/v9"
)

var RedisDb *redis.Client

func NewRedisCoontext() error {

	conf := pkg.Config{}.LOAD()
	client := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddress,
		Password: "",
		DB:       0,
	})
	RedisDb = client

	return checkRedisHealth()
}

func checkRedisHealth() error {
	ctx := context.Background()
	res, err := RedisDb.Ping(ctx).Result()
	fmt.Println("Redis Response for PING : ", res)
	return err
}
