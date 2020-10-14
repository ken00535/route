package router

const (
	errTopicNotFound int = 1
)

// RouterError is error type of router
type RouterError struct {
	message string
	errno   int
}

func (e RouterError) Error() string {
	return e.message
}

// IsTopicNotFound is topic error or not
func (e RouterError) IsTopicNotFound() bool {
	return e.errno == errTopicNotFound
}
