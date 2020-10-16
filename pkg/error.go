package router

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

// IsTopicNotFound is topic error or not
func (e Error) IsTopicNotFound() bool {
	return e.errno == errTopicNotFound
}
