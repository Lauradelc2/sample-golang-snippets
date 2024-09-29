package logic

import (
	
	"context"
	"greet/internal/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGreet(t *testing.T) {
	req := &types.Request{
		Name: "test",
	}
	gl := NewGreetLogic(context.Background(), nil)
	res, _ := gl.Greet(req)

	assert.Equal(t, res.Message, "hello test")
}
