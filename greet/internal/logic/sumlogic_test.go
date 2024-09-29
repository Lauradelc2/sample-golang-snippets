package logic

import (
	
	"context"
	"greet/internal/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	val := &types.ArithmeticReq{
		O1: 1,
		O2: 2,
	}
	sl := NewSumLogic(context.Background(), nil)
	res, _ := sl.Sum(val)

	assert.Equal(t, res.Val, 3)
}
