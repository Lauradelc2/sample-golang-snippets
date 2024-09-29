package logic

import (
	"context"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubLogic {
	return &SubLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubLogic) Sub(req *types.ArithmeticReq) (resp *types.Result, err error) {
	return &types.Result{
		Val: req.O1 - req.O2,
	}, nil
}
