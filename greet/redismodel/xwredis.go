package redismodel

import (
	//"database/sql"
	"fmt"
	"github.com/go-redis/redis/v9"
)

type RedisBasicModel struct {
	conn *redis.Conn
}

var RDB RedisBasicModel
func InitRedis() *redis.Client{
	RDB.(*redis.Client)= redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return RDB
}