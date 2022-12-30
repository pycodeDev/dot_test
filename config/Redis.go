package config

import (
	"time"

	"github.com/go-redis/redis/v8"
)

func Redislocal() *redis.Client {
	k := GetMyConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:        k.REDIS.REDIS_HOST_LOCAL + ":" + k.REDIS.REDIS_PORT_LOCAL,
		Password:    k.REDIS.REDIS_PASS_LOCAL, // no password set
		DB:          0,                        // use default DB
		PoolTimeout: time.Second * 1,
		IdleTimeout: time.Second * 1,
	})
	return rdb
}
