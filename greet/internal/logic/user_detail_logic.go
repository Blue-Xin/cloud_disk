package logic

import (
	"context"
	"errors"
	"zero/greet/models"

	"zero/greet/internal/svc"
	"zero/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailReply, err error) {
	// todo: add your logic here and delete this line
	resp = &types.UserDetailReply{}
	ub := new(models.UserBasic)
	get, err := models.Engine.Where("identity=?", req.Identity).Get(ub)
	if err != nil {
		return nil, err
	}
	if !get {
		return nil, errors.New("user not found")
	}
	resp.Name = ub.Name
	resp.Email = ub.Email
	return
}
