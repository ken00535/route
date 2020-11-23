package route

import "fmt"

// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Context)

// HandlersChain defines a HandlerFunc array.
type HandlersChain []HandlerFunc

// New a Router
func New() Router {
	r := Router{}
	r.nodes = make(map[string]*node)
	return r
}

// Router process message
type Router struct {
	routeRule     func(interface{}) string
	nodes         map[string]*node
	useMiddleware HandlersChain
}

// Use attach middleware to every route
func (r *Router) Use(m ...HandlerFunc) {
	r.useMiddleware = append(r.useMiddleware, m...)
}

// Add add middleware
func (r *Router) Add(topic string, m ...HandlerFunc) {
	// * is for default topic
	if topic == "*" {
		topic = defaultTopic
	}
	if _, exist := r.nodes[topic]; !exist {
		r.nodes[topic] = &node{}
	}
	n := r.nodes[topic]
	m = append(r.useMiddleware, m...)
	n.MergeHandler(m...)
}

// Run router
func (r *Router) Run(message interface{}) error {
	var topic string
	if r.routeRule != nil {
		topic = r.routeRule(message)
	} else {
		topic = defaultTopic
	}
	// do nothing if no topic
	if topic == "" {
		return nil
	}
	if n, exist := r.nodes[topic]; exist {
		c := Context{}
		c.reset()
		c.handlers = n.handlers
		c.Message = message
		c.Next()
		return c.Errors.Last()
	}
	return Error{
		message: fmt.Sprintf("topic %s is not registered.", topic),
		errno:   errTopicNotFound,
	}
}

// SetRouteRule set rule
func (r *Router) SetRouteRule(callback func(interface{}) string) {
	r.routeRule = callback
}
