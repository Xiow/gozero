package model

import "gorm.io/gorm"

//type (
//	gormBasicModel interface {
//		Insert()string
//	}
//)
type (
	gormBasicModel interface {
		Insert()string
	}
)
type (
	GormBasicModel interface {
		gormBasicModel
	}
	customGormBasicModel struct {
		*defaultGormBasicModel
	}
)

func (c customGormBasicModel) Insert() string {
	return "Hello,Gorm"
}
func NewGormBasicModel(gormClient *gorm.DB, err error) GormBasicModel {
	return &customGormBasicModel{
		defaultGormBasicModel: newGormBasicModel(gormClient),
	}
}
func newGormBasicModel(client *gorm.DB )*defaultGormBasicModel{
	return &defaultGormBasicModel{
		conn:  client,
	}
}
