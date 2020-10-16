package router

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	var actual string
	middle := func(next Handler) Handler {
		return func(res interface{}) error {
			actual = res.(string) + " world"
			return next(res)
		}
	}
	msg := "hello"
	expect := "hello world"
	processor := New()
	processor.Use(middle)
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
	msg := "topic1"
	switchRule := func(message interface{}) string {
		return message.(string)
	}
	expect := []string{"world"}
	processor := New()
	processor.SetRouteRule(switchRule)
	processor.UseTopic("topic1", mid1)
	processor.UseTopic("topic2", mid2)
	processor.Run(msg)
	assert.Equal(t, expect, actual)
}

func TestReturnError(t *testing.T) {
	mid := func(next Handler) Handler {
		return func(res interface{}) error {
			return errors.New("this is an error")
		}
	}
	msg := "topic1"
	switchRule := func(message interface{}) string {
		return message.(string)
	}
	expect := errors.New("this is an error")
	processor := New()
	processor.SetRouteRule(switchRule)
	processor.UseTopic("topic1", mid)
	actual := processor.Run(msg)
	assert.Equal(t, expect, actual)
}

func TestNoTopicError(t *testing.T) {
	mid := func(next Handler) Handler {
		return func(res interface{}) error {
			return errors.New("this is an error")
		}
	}
	msg := "topic"
	switchRule := func(message interface{}) string {
		return message.(string)
	}
	processor := New()
	processor.SetRouteRule(switchRule)
	processor.UseTopic("topic1", mid)
	actual := processor.Run(msg).(Error).IsTopicNotFound()
	assert.Equal(t, true, actual)
}
