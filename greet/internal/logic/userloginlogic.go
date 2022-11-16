package logic

import (
	"context"
	"greet/internal/svc"
	"greet/internal/types"
	"greet/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	//user :=model.LoginBasic{"dawei","123456"}
	loginbasic,err:=model.LoginBasicModel.FindOne(l.svcCtx.LoginBasicModel,l.ctx,"xw")
	//str:=l.svcCtx.RedisModel.Insert()
	//fmt.Println(str)
	//gorm
	//fmt.Println(l.svcCtx.GormModel.Insert())
	return &types.LoginResponse{
		Name: loginbasic.Name,
		Password: loginbasic.Password,
		Token: "123",
	},nil

}
