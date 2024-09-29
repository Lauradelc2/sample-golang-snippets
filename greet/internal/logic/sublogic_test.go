package logic

import (
	
	"context"
	"greet/internal/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSub(t *testing.T) {
	val := &types.ArithmeticReq{
		O1: 2,
		O2: 1,
	}
	sl := NewSubLogic(context.Background(), nil)
	res, _ := sl.Sub(val)

	assert.Equal(t, res.Val, float32(1))
}
