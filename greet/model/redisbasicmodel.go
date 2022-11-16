package model

import (
	"context"
	"github.com/go-redis/redis/v9"
)

type (
customRedisBasicModel struct {
*defaultRedisBasicModel
}
RedisBasicModel interface {
redisBasicModel
}
)

func (c customRedisBasicModel) Insert()string{
	return "Hello,Redis"
}
func NewRedisModel(redisClient *redis.Client) RedisBasicModel{
	//return RedisBasicModel(redisClient)
	return &customRedisBasicModel{
		defaultRedisBasicModel:newRedisBasicModel(redisClient),
	}
}
func newRedisBasicModel(client *redis.Client) *defaultRedisBasicModel{
	return &defaultRedisBasicModel{
		conn:  client,
	}
}
func InitRedisClient()*redis.Client{
	rdb:=redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil
	}
	return rdb
}