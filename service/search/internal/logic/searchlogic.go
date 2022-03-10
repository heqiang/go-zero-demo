package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"go-zero-demo/service/search/internal/svc"
	"go-zero-demo/service/search/internal/types"
	"go-zero-demo/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchLogic {
	return SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req types.SearchReq) (resp *types.SearchReplay, err error) {
	userIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId")))
	logx.Infof("userId: %s", userIdNumber)
	userId, err := userIdNumber.Int64()
	if err != nil {
		return nil, err
	}

	// 使用user rpc
	_, err = l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.IdReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.SearchReplay{
		Name:  req.Name,
		Count: 100,
	}, nil
}
