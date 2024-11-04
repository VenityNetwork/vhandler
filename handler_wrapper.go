package vhandler

type handlerWrapper[T any] struct {
	priority Priority
	h        T
}
