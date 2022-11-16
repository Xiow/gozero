package test

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	//"greet/helper/Xwviper"
	"gorm.io/gorm"
	"testing"
)

func TestGormConnect(t *testing.T){
	//"mysql", "root:123456@(localhost)/gormtest?charset=utf8mb4&parseTime=True&loc=Local"

	db,err:=gorm.Open(mysql.Open("root:123456@(localhost)/gormtest?charset=utf8mb4&parseTime=True&loc=Local"),&gorm.Config{})
	if err!=nil{
		t.Log(err)
	}
	t.Log(db)
}
func TestRedisConn(t *testing.T){
	rdb:=redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		t.Log(err)
	}else{
		t.Log("success")
	}
}
