package vhandler

type subHandler struct {
	handlers map[priorityId][]Handler
}

func newSubHandler() *subHandler {
	return &subHandler{handlers: map[priorityId][]Handler{}}
}

func (s *subHandler) handle(f func(Handler)) {
	for _, id := range priorityOrder {
		for _, h := range s.handlers[id] {
			f(h)
		}
	}
}

func (s *subHandler) add(p priorityId, h Handler) {
	s.handlers[p] = append(s.handlers[p], h)
}
