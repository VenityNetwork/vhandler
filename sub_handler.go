package vhandler

import "github.com/venitynetwork/vhandler/priority"

type subHandler struct {
	handlers      map[priority.Priority][]Handler
	priorityOrder []priority.Priority
}

func newSubHandler() *subHandler {
	return &subHandler{handlers: map[priority.Priority][]Handler{}}
}

func (s *subHandler) handle(f func(Handler)) {
	for _, id := range s.priorityOrder {
		for _, h := range s.handlers[id] {
			f(h)
		}
	}
}

func (s *subHandler) add(p priority.Priority, h Handler) {
	s.handlers[p] = append(s.handlers[p], h)

	s.priorityOrder = []priority.Priority{}
	for _, pid := range priority.Order {
		if _, ok := s.handlers[pid]; ok {
			s.priorityOrder = append(s.priorityOrder, pid)
		}
	}
}
