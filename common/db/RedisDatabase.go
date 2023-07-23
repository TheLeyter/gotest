package db

import (
	"fmt"
	"gotest/common/envs"

	"github.com/redis/go-redis/v9"
)

type RedisDatabase struct {
	Config *envs.RedisConfig
	Redis  *redis.Client
}

func NewRedisDatabse(Config *envs.RedisConfig) *RedisDatabase {
	return &RedisDatabase{Config: Config}
}

func (redisDb *RedisDatabase) Connect() {
	host := redisDb.Config.Host + ":" + redisDb.Config.Port

	options := redis.Options{Addr: host, Username: redisDb.Config.User, Password: redisDb.Config.Password}

	redisDb.Redis = redis.NewClient(&options)

	fmt.Println("-------" + redisDb.Redis.String() + "--------")
}
