package logic

import (
	
	"context"
	"greet/internal/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMul(t *testing.T) {
	val := &types.ArithmeticReq{
		O1: 2,
		O2: 1,
	}
	ml := NewMulLogic(context.Background(), nil)
	res, _ := ml.Mul(val)

	assert.Equal(t, res.Val, 2)
}
