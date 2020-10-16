# router

router provides a middleware pattern

## Usage

```go
	mid1 := func(next Handler) Handler {
		return func(res interface{}) error {
			fmt.Println("this is topic 1 middleware")
			return next(res)
		}
	}
	mid2 := func(next Handler) Handler {
		return func(res interface{}) error {
			fmt.Println("this is topic 2 middleware")
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
    // -> print "this is topic 1 middleware"
```

## Error Handling

you can handle error by

```go
    processor.UseTopic("topic1", mid)
    err := processor.Run(msg)
    if err != nil {
        if err.(router.Error).IsTopicNotFound() {
            // do something...
        } else {
            return err
        }
    }
```
