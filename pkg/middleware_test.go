package gonion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	var actual string
	mid := func(next Handler) Handler {
		return func(res interface{}) error {
			actual = res.(string) + " world"
			return next(res)
		}
	}
	msg := "hello"
	expect := "hello world"
	processor := New()
	processor.Use(mid)
	processor.Run(msg)
	assert.Equal(t, expect, actual)
}
