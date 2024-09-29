package logic

import (
	"context"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MulLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMulLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MulLogic {
	return &MulLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MulLogic) Mul(req *types.ArithmeticReq) (resp *types.Result, err error) {
	return &types.Result{
		Val: req.O1 * req.O2,
	}, nil
}
