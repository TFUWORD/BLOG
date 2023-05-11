package core

import (
	"context"
	"fmt"
	"time"
	"web/global"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)



func ConnectRedis() *redis.Client {
	return ConnectRedisDB(0)
}

func ConnectRedisDB(db int) *redis.Client {
	redisConf := global.Config.Redis
  rdb := redis.NewClient(&redis.Options{
    Addr:     redisConf.Addr(),
    Password: redisConf.Password, // no password set
    DB:       0,       // use default DB
    PoolSize: redisConf.PoolSize,     // 连接池大小
  })
	fmt.Println(rdb)
  _, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
  defer cancel()
  _, err := rdb.Ping().Result()
  if err != nil {
    logrus.Error("redis连接失败 %s", redisConf.Addr())
    return nil
  }
	return rdb
}
