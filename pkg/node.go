package route

type node struct {
	handlers HandlersChain
}

func (n *node) MergeHandler(h ...HandlerFunc) {
	n.handlers = append(n.handlers, h...)
}
