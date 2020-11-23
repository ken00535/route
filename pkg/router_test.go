package route

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	var actual string
	handler := func(c *Context) {
		actual = c.Message.(string) + " world"
		c.Next()
	}
	msg := "hello"
	expect := "hello world"
	router := New()
	router.Use("*", handler)
	router.Run(msg)
	assert.Equal(t, expect, actual)
}

func TestTwoTopic(t *testing.T) {
	var actual []string
	mid1 := func(c *Context) {
		actual = append(actual, "world")
		c.Next()
	}
	mid2 := func(c *Context) {
		actual = append(actual, "hello")
		c.Next()
	}
	msg := "topic1"
	switchRule := func(message interface{}) string {
		return message.(string)
	}
	expect := []string{"world"}
	router := New()
	router.SetRouteRule(switchRule)
	router.Use("topic1", mid1)
	router.Use("topic2", mid2)
	router.Run(msg)
	assert.Equal(t, expect, actual)
}

func TestReturnError(t *testing.T) {
	mid := func(c *Context) {
		c.Error(errors.New("this is an error"))
	}
	msg := "topic1"
	switchRule := func(message interface{}) string {
		return message.(string)
	}
	expect := &Error{message: "this is an error"}
	router := New()
	router.SetRouteRule(switchRule)
	router.Use("topic1", mid)
	actual := router.Run(msg)
	assert.Equal(t, expect, actual)
}

func TestNoTopicError(t *testing.T) {
	mid := func(c *Context) {
		c.Error(errors.New("this is an error"))
	}
	msg := "topic"
	switchRule := func(message interface{}) string {
		return message.(string)
	}
	router := New()
	router.SetRouteRule(switchRule)
	router.Use("topic1", mid)
	actual := router.Run(msg).(Error).IsTopicNotFound()
	assert.Equal(t, true, actual)
}
