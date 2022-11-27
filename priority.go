package vhandler

type priorityId uint8

const (
	PriorityLowest priorityId = iota
	PriorityLow
	PriorityNormal
	PriorityHigh
	PriorityHighest
	PriorityMonitor
)

var priorityOrder = []priorityId{PriorityLowest, PriorityLow, PriorityNormal, PriorityHigh, PriorityHighest, PriorityMonitor}
