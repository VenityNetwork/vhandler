package priority

type Priority uint8

const (
	Lowest Priority = iota
	Low
	Normal
	High
	Highest
	Monitor
)

var Order = []Priority{Lowest, Low, Normal, High, Highest, Monitor}
