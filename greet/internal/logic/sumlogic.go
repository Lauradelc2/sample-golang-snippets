package logic

import (
	"context"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SumLogic {
	return &SumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SumLogic) Sum(req *types.ArithmeticReq) (resp *types.Result, err error) {
	return &types.Result{
		Val: req.O1 + req.O2,
	}, nil
}
