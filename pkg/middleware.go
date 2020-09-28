package gonion

// Handler is query handler
type Handler func(interface{}) error

// Middleware is public middleware
type Middleware func(Handler) Handler

// New a processor
func New() Processor {
	return Processor{}
}

// Processor process message
type Processor struct {
	handler         Handler
	middlewareChain []Middleware
}

// Use add middleware
func (r *Processor) Use(m ...Middleware) {
	r.middlewareChain = append(r.middlewareChain, m...)
}

// Run router
func (r *Processor) Run(message interface{}) error {
	r.handler = func(in interface{}) error {
		return nil
	}
	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		r.handler = r.middlewareChain[i](r.handler)
	}
	return r.handler(message)
}
