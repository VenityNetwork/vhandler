package vhandler

import "sort"

func sortHandlers[T any](handlers []*handlerWrapper[T]) {
	sort.Slice(handlers, func(i, j int) bool {
		return handlers[i].priority < handlers[j].priority
	})
}
