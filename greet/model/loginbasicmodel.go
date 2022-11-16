package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LoginBasicModel = (*customLoginBasicModel)(nil)

type (
	// LoginBasicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLoginBasicModel.
	LoginBasicModel interface {
		loginBasicModel
	}

	customLoginBasicModel struct {
		*defaultLoginBasicModel
	}
)
type (
	redisBasicModel interface {
		Insert()string
	}
)
// NewLoginBasicModel returns a model for the database table.
func NewLoginBasicModel(conn sqlx.SqlConn) LoginBasicModel {
	return &customLoginBasicModel{
		defaultLoginBasicModel: newLoginBasicModel(conn),
	}
}



