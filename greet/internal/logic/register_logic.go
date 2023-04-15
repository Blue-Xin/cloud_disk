package logic

import (
	"context"
	"errors"
	"time"
	"zero/greet/define"
	"zero/greet/helper"
	"zero/greet/models"

	"zero/greet/internal/svc"
	"zero/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.MailCodeRequest) (resp *types.MailCodeReply, err error) {
	//该邮箱不存在
	count, err := models.Engine.Where("email=?", req.Email).Count(new(models.UserBasic))
	if err != nil {
		return nil, err
	}
	if count > 0 {
		err = errors.New("邮箱已被注册")
	}

	code := helper.RandCode()
	//存储验证码
	models.RDB.Set(l.ctx, req.Email, code, time.Second*time.Duration(define.CodeExpire))

	err = helper.MailSendCode(req.Email, code)
	if err != nil {
		return nil, err
	}
	return
}
