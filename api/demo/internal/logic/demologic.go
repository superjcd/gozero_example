package logic

import (
	"context"

	"github.com/superjcd/go-zero-templates/api/demo/internal/svc"
	"github.com/superjcd/go-zero-templates/api/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DemoLogic {
	return &DemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DemoLogic) Demo(req *types.Request) (resp *types.Response, err error) {
	resp = new(types.Response)
	resp.Message = req.Name
	return
}
