# route

route provides a route mechanism, you can route your message easily.

for example, if you have to handle your customized protocol (a.k.a not http), you can use it!

this project is inspired by [Gin](https://github.com/gin-gonic/gin) and use source code of Gin

## Install

```bash
go get github.com/ken00535/router
```

## Usage

```go
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
router.Add("topic1", mid1)
router.Add("topic2", mid2)
router.Run(msg)
```

## Custom Middleware

```go
var actual string
addHandler := func(c *Context) {
    actual += " world"
    c.Next()
}
useHandler := func(c *Context) {
    actual = "hello lovely"
    c.Next()
}
expect := "hello lovely world"
router := New()
router.Use(useHandler)
router.Add("*", addHandler)
router.Run("")
```

## Error Handling

you can handle error by

```go
mid := func(c *Context) {
    c.Error(errors.New("this is an error"))
}
msg := "topic"
switchRule := func(message interface{}) string {
    return message.(string)
}
router := New()
router.SetRouteRule(switchRule)
router.Add("topic1", mid)
actual := router.Run(msg).(Error).IsTopicNotFound()
```
