package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"go-zero-demo/common/errorx"
	"go-zero-demo/service/user/model"
	"strings"
	"time"

	"go-zero-demo/service/user/api/internal/svc"
	"go-zero-demo/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (resp *types.LoginReply, err error) {
	if len(strings.TrimSpace(req.Username))==0||len(strings.TrimSpace(req.Password))==0{
		return nil,errorx.NewDefaultError("用户名或者密码不能为空")
	}
	userInfo,err:= l.svcCtx.UserModel.FindOneByNumber(req.Username)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil,errorx.NewDefaultError("用户不存在")
	default:
		return nil,err
	}

	if userInfo.Password!=req.Password{
		return nil,errorx.NewDefaultError("密码错误")
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	sercetKey := l.svcCtx.Config.Auth.AccessSecret
	token, err := l.getJwtToken(sercetKey, now, accessExpire, userInfo.Id)
	if err != nil {
		return nil, err
	}

	return &types.LoginReply{
		Id :          userInfo.Id,
		Name:         userInfo.Name,
		Gender:       userInfo.Gender,
		AccessToken:  token,
		AccessExpire: now+accessExpire,
		RefreshAfter: now + accessExpire/2,
	},nil
	
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
