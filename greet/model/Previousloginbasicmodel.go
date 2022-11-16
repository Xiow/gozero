package model
//
//import (
//	//"github.com/lib/pq"
//	"github.com/go-redis/redis/v9"
//	"github.com/zeromicro/go-zero/core/stores/sqlx"
//	"gorm.io/gorm"
//)
//
//var _ LoginBasicModel = (*customLoginBasicModel)(nil)
//
//type (
//	// LoginBasicModel is an interface to be customized, add more methods here,
//	// and implement the added methods in customLoginBasicModel.
//	LoginBasicModel interface {
//		loginBasicModel
//	}
//
//	customLoginBasicModel struct {
//		*defaultLoginBasicModel
//	}
//	customRedisBasicModel struct {
//		*defaultRedisBasicModel
//	}
//	RedisBasicModel interface {
//		redisBasicModel
//	}
//	//gormModel
//	GormBasicModel interface {
//		gormBasicModel
//	}
//	customGormBasicModel struct {
//		*defaultGormBasicModel
//	}
//)
//
//func (c customGormBasicModel) Insert() string {
//	return "Hello,Gorm"
//}
//
//func (c customRedisBasicModel) Insert()string{
//	return "Hello,Redis"
//}
//
//type (
//	redisBasicModel interface {
//		Insert()string
//	}
//)
//
//type (
//	gormBasicModel interface {
//		Insert()string
//	}
//)
//// NewLoginBasicModel returns a model for the database table.
//func NewLoginBasicModel(conn sqlx.SqlConn) LoginBasicModel {
//	return &customLoginBasicModel{
//		defaultLoginBasicModel: newLoginBasicModel(conn),
//	}
//}
//func NewRedisModel(redisClient *redis.Client) RedisBasicModel{
//	//return RedisBasicModel(redisClient)
//	return &customRedisBasicModel{
//		defaultRedisBasicModel:newRedisBasicModel(redisClient),
//	}
//}
//func NewGormBasicModel(gormClient *gorm.DB,err error) GormBasicModel{
//	return &customGormBasicModel{
//		defaultGormBasicModel:newGormBasicModel(gormClient),
//	}
//}
//
//
//
