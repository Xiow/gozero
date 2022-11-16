package svc

import (
	"github.com/go-redis/redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"greet/internal/config"
	"greet/internal/middleware"
	"greet/model"
)

type ServiceContext struct {
	Config             config.Config
	LoginBasicModel    model.LoginBasicModel
	RegisterBasicModel model.RegisterBasicModel
	RedisModel         model.RedisBasicModel
	GormModel          model.GormBasicModel
	XwMiddleware       rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//LoginBasicModel: model.NewLoginBasicModel(sqlx.NewMysql(c.MysqlAddr)),
		XwMiddleware:	middleware.NewXwMiddleware().Handle,
		RegisterBasicModel:model.NewRegisterBasicModel(sqlx.NewMysql(c.MysqlAddr)),
		RedisModel:model.NewRedisModel(redis.NewClient(&redis.Options{Addr: "localhost:6379"})),
		GormModel:model.NewGormBasicModel(gorm.Open(mysql.Open("root:123456@(localhost)/gormtest?charset=utf8mb4&parseTime=True&loc=Local"),&gorm.Config{})),
	}
}
