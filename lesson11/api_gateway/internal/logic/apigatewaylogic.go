// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"api_gateway/internal/svc"
	"api_gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Api_gatewayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApi_gatewayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Api_gatewayLogic {
	return &Api_gatewayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Api_gatewayLogic) Api_gateway(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
