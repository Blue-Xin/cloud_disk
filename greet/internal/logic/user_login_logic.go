package logic

import (
	"context"
	"errors"
	"fmt"
	"zero/greet/helper"
	"zero/greet/models"

	"zero/greet/internal/svc"
	"zero/greet/internal/types"

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

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line
	//从数据库查询当前用户
	user := new(models.UserBasic)
	get, err := models.Engine.Where("name=? and password = ?", req.Name, helper.Md5(req.Password)).Get(user)
	md5 := helper.Md5(req.Password)
	fmt.Println(md5)
	if err != nil {
		return nil, err
	}
	if !get {
		return nil, errors.New("用户名或密码错误")
	}
	//查询到返回用户信息
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name)
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginReply)
	resp.Token = token
	//生成token

	return
}
