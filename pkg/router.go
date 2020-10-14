package router

import "fmt"

// Handler is query handler
type Handler func(interface{}) error

// Middleware is public middleware
type Middleware func(Handler) Handler

// New a Router
func New() Router {
	r := Router{}
	r.middlewareChains = make(map[string][]Middleware)
	r.handlers = make(map[string]Handler)
	return r
}

// Router process message
type Router struct {
	routeRule        func(interface{}) string
	middlewareChains map[string][]Middleware
	handlers         map[string]Handler
}

// UseTopic add middleware
func (r *Router) UseTopic(topic string, m ...Middleware) {
	r.middlewareChains[topic] = append(r.middlewareChains[topic], m...)
	r.handlers[topic] = func(interface{}) error {
		return nil
	}
}

// Use add middleware
func (r *Router) Use(m ...Middleware) {
	r.middlewareChains[defaultTopic] = append(r.middlewareChains[defaultTopic], m...)
	r.handlers[defaultTopic] = func(interface{}) error {
		return nil
	}
}

// Run router
func (r *Router) Run(message interface{}) error {
	for k, middlewareChain := range r.middlewareChains {
		for i := len(middlewareChain) - 1; i >= 0; i-- {
			r.handlers[k] = middlewareChain[i](r.handlers[k])
		}
	}
	var topic string
	if r.routeRule != nil {
		topic = r.routeRule(message)
	} else {
		topic = defaultTopic
	}
	if r.handlers[topic] != nil {
		return r.handlers[topic](message)
	}
	return RouterError{
		message: fmt.Sprintf("topic %s is not registered.", topic),
		errno:   errTopicNotFound,
	}
}

// SetRouteRule set rule
func (r *Router) SetRouteRule(callback func(interface{}) string) {
	r.routeRule = callback
}
