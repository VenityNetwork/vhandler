package vhandler

import (
	"errors"
	"github.com/venitynetwork/vhandler/priority"
	"sync"
)

type subHandler struct {
	mu sync.Mutex

	handlers map[priority.Priority][]Handler
}

func newSubHandler() *subHandler {
	return &subHandler{handlers: map[priority.Priority][]Handler{}}
}

func (s *subHandler) handle(f func(Handler)) {
	for _, id := range priority.Order {
		for _, h := range s.handlers[id] {
			f(h)
		}
	}
}

func (s *subHandler) add(p priority.Priority, h Handler) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.handlers[p] = append(s.handlers[p], h)
}

func (s *subHandler) remove(h Handler) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	deleted := false

	for i, handlers := range s.handlers {
		var newHandlers []Handler
		for _, handler := range handlers {
			if handler == h {
				deleted = true
				continue
			}
			newHandlers = append(newHandlers, handler)
		}
		if len(newHandlers) == 0 {
			delete(s.handlers, i)
			continue
		}
		s.handlers[i] = newHandlers
	}

	if deleted {
		return nil
	}

	return errors.New("failed to delete handler: handler not found")
}
