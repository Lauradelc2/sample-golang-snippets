package logic

import (
	
	"context"
	"greet/internal/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiv(t *testing.T) {
	val := &types.ArithmeticReq{
		O1: 2,
		O2: 1,
	}
	ml := NewDivLogic(context.Background(), nil)
	res, _ := ml.Div(val)

	assert.Equal(t, res.Val, float32(2))
}
