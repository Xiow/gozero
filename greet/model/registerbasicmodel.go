package model

import (
	//"github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RegisterBasicModel = (*customRegisterBasicModel)(nil)

type (
	// RegisterBasicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRegisterBasicModel.
	RegisterBasicModel interface {
		registerBasicModel
	}

	customRegisterBasicModel struct {
		*defaultRegisterBasicModel
	}
)

// NewRegisterBasicModel returns a model for the database table.
func NewRegisterBasicModel(conn sqlx.SqlConn) RegisterBasicModel {
	return &customRegisterBasicModel{
		defaultRegisterBasicModel: newRegisterBasicModel(conn),
	}
}
