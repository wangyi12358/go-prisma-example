package logic

import (
	"context"
	"fmt"
	"go-zero-micro/service/user/model/sys_user_model"

	"go-zero-micro/service/user/rpc/internal/svc"
	"go-zero-micro/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserLogic {
	return &CheckUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUserLogic) CheckUser(in *user.CheckUserRequest) (*user.CheckUserResponse, error) {
	sysUser, err := sys_user_model.FindOneByLogin(in.Username, in.Password)
	if err != nil {
		fmt.Printf("check user error: %s", err.Error())
		return &user.CheckUserResponse{Success: false}, nil
	}
	return &user.CheckUserResponse{
		Success: sysUser != nil,
	}, nil
}
