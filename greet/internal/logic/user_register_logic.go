package logic

import (
	"context"
	"errors"
	"log"
	"zero/greet/helper"
	"zero/greet/internal/svc"
	"zero/greet/internal/types"
	"zero/greet/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {
	//判断code是否一致
	code, err := models.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("该邮箱验证码为空")
	}
	if code != req.Code {
		err = errors.New("验证码错误")
		return
	}
	//判断用户名是否存在
	count, err := models.Engine.Where("name=?", req.Name).Count(new(models.UserBasic))
	if count > 0 {
		err = errors.New("用户名已存在")
		return
	}
	//数据库入库
	user := &models.UserBasic{
		Identity: helper.UUID(),
		Name:     req.Name,
		Password: helper.Md5(req.Password),
		Email:    req.Email,
	}
	insert, err := models.Engine.Insert(user)
	if err != nil {
		return nil, err
	}
	log.Println("insert user row:", insert)

	return
}
