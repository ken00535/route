package gonion

import (
	"errors"
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

func TestTwoTopic(t *testing.T) {
	var actual []string
	mid1 := func(next Handler) Handler {
		return func(res interface{}) error {
			actual = append(actual, "world")
			return next(res)
		}
	}
	mid2 := func(next Handler) Handler {
		return func(res interface{}) error {
			actual = append(actual, "hello")
			return next(res)
		}
	}
	msg := ""
	expect := []string{"hello", "world"}
	processor := New()
	processor.UseTopic("topic1", mid1)
	processor.UseTopic("topic2", mid2)
	processor.Run(msg)
	assert.Equal(t, expect, actual)
}

func TestReturnError(t *testing.T) {
	mid1 := func(next Handler) Handler {
		return func(res interface{}) error {
			return errors.New("this is an error")
		}
	}
	mid2 := func(next Handler) Handler {
		return func(res interface{}) error {
			return errors.New("this is another error")
		}
	}
	msg := ""
	expect := []error{
		errors.New("this is an error"),
		errors.New("this is another error"),
	}
	processor := New()
	processor.UseTopic("topic1", mid1)
	processor.UseTopic("topic2", mid2)
	actual := processor.Run(msg)
	assert.Equal(t, expect, actual)
}
