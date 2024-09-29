package logic

import (
	"context"
	"errors"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DivLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDivLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DivLogic {
	return &DivLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DivLogic) Div(req *types.ArithmeticReq) (resp *types.Result, err error) {
	if req.O2 == 0 {
		return nil, errors.New("division by zero not allowed")
	}

	return &types.Result{
		Val: req.O1 / req.O2,
	}, nil
}
