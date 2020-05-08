package redis

import (
	"fmt"
	"time"

	GoRedis "github.com/go-redis/redis"
	"github.com/kxiaong/blog/conf"
)

var GoRedisClient *GoRedis.Client

func Init() {
	GoRedisClient = GoRedis.NewClient(&GoRedis.Options{
		Addr:         conf.C.Redis.Addr,
		Password:     conf.C.Redis.Password,
		MaxRetries:   3,
		DialTimeout:  2 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolTimeout:  5 * time.Second,
		DB:           conf.C.Redis.DB,
		PoolSize:     conf.C.Redis.PoolSize,
		IdleTimeout:  15 * time.Second,
	})

	if pong, err := GoRedisClient.Ping().Result(); err != nil || pong != "PONG" {
		panic(fmt.Sprintf("Ping redis failed: %s", err.Error()))
	}
}

func Lock(key string, timeout time.Duration) (bool, error) {
	return GoRedisClient.SetNX(key, 1, timeout).Result()
}

func UnLock(key string) (int64, error) {
	return GoRedisClient.Del(key).Result()
}
