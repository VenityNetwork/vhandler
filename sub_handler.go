package vhandler

type subHandler struct {
	handlers      map[priorityId][]Handler
	priorityOrder []priorityId
}

func newSubHandler() *subHandler {
	return &subHandler{handlers: map[priorityId][]Handler{}}
}

func (s *subHandler) handle(f func(Handler)) {
	for _, id := range s.priorityOrder {
		for _, h := range s.handlers[id] {
			f(h)
		}
	}
}

func (s *subHandler) add(p priorityId, h Handler) {
	s.priorityOrder = []priorityId{}
	for _, pid := range priorityOrder {
		if _, ok := s.handlers[pid]; ok {
			s.priorityOrder = append(s.priorityOrder, pid)
		}
	}
	s.handlers[p] = append(s.handlers[p], h)
}
