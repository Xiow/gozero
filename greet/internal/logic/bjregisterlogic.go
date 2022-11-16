package logic

import (
	"context"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BjRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBjRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BjRegisterLogic {
	return &BjRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BjRegisterLogic) BjRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
