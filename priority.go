package vhandler

// Priority represents the priority of a handler. The priority of a handler determines the order in which
// handlers are called. Handlers with a higher priority are called before handlers with a lower priority.
type Priority int

const (
	PriorityLowest  Priority = 100
	PriorityLow              = 200
	PriorityNormal           = 300
	PriorityHigh             = 400
	PriorityHighest          = 500
)
