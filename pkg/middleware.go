package gonion

// Handler is query handler
type Handler func(interface{}) error

// Middleware is public middleware
type Middleware func(Handler) Handler

// New a processor
func New() Processor {
	p := Processor{}
	p.callbacks = make(map[string]Handler)
	p.middlewareChains = make(map[string][]Middleware)
	p.callbackChannels = make(map[string]chan error)
	return p
}

// Processor process message
type Processor struct {
	callbacks        map[string]Handler
	middlewareChains map[string][]Middleware
	callbackChannels map[string]chan error
}

// UseTopic add middleware
func (r *Processor) UseTopic(topic string, m ...Middleware) {
	r.middlewareChains[topic] = append(r.middlewareChains[topic], m...)
}

// Use add middleware
func (r *Processor) Use(m ...Middleware) {
	r.middlewareChains[defaultTopic] = append(r.middlewareChains[defaultTopic], m...)
}

// Run router
func (r *Processor) Run(message interface{}) []error {
	var ret []error
	for k, middlewareChain := range r.middlewareChains {
		r.callbacks[k] = func(in interface{}) error {
			return nil
		}
		for i := len(middlewareChain) - 1; i >= 0; i-- {
			r.callbacks[k] = middlewareChain[i](r.callbacks[k])
		}
		r.callbackChannels[k] = make(chan error, 1)
		go func(topic string) {
			err := r.callbacks[topic](message)
			r.callbackChannels[topic] <- err
		}(k)
	}
	for k := range r.callbackChannels {
		err := <-r.callbackChannels[k]
		if err != nil {
			ret = append(ret, err)
		}
	}
	return ret
}
