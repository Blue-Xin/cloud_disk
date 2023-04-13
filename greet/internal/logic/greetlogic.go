package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"zero/greet/models"

	"zero/greet/internal/svc"
	"zero/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetLogic {
	return &GreetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetLogic) Greet(req *types.Request) (resp *types.Response, err error) {
	resp = new(types.Response)

	data := make([]*models.UserBasic, 0)
	//获取用户列表
	err = models.Engine.Find(&data)
	if err != nil {
		log.Printf("Get UserBasic Error:", err)
	}
	marshal, err := json.Marshal(data)
	if err != nil {
		log.Printf("Json UserBasic Error:", err)
	}
	b := new(bytes.Buffer)
	err = json.Indent(b, marshal, "", " ")
	if err != nil {
		log.Printf("Json Indent Error:", err)

	}
	fmt.Println(b)
	resp.Message = b.String()
	return
}
