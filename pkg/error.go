package route

const (
	errTopicNotFound int = 1
)

// Error is error type of router
type Error struct {
	message string
	errno   int
}

func (e Error) Error() string {
	return e.message
}

type errorMsgs []*Error

// Last returns the last error in the slice. It returns nil if the array is empty.
// Shortcut for errors[len(errors)-1].
func (a errorMsgs) Last() *Error {
	if length := len(a); length > 0 {
		return a[length-1]
	}
	return nil
}

// Errors returns an array will all the error messages.
// Example:
// 		c.Error(errors.New("first"))
// 		c.Error(errors.New("second"))
// 		c.Error(errors.New("third"))
// 		c.Errors.Errors() // == []string{"first", "second", "third"}
func (a errorMsgs) Errors() []string {
	if len(a) == 0 {
		return nil
	}
	errorStrings := make([]string, len(a))
	for i, err := range a {
		errorStrings[i] = err.Error()
	}
	return errorStrings
}

// IsTopicNotFound is topic error or not
func (e Error) IsTopicNotFound() bool {
	return e.errno == errTopicNotFound
}
