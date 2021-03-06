package logic

import (
	"context"

	"go-zero-demo/service/user/rpc/internal/svc"
	"go-zero-demo/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdReq) (*user.UserInfoReplay, error) {
	one ,err:=l.svcCtx.UserModel.FindOne(in.Id)
	if err!=nil{
		return nil,err
	}
	return &user.UserInfoReplay{
		Id: one.Id,
		Name: one.Name,
		Number: one.Number,
		Gender: one.Gender,
	}, nil
}
