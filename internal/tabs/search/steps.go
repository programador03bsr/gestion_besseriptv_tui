package search

type step int

const (
	stepTyping step = iota
	stepSelecting
	stepProcessing
	stepDone
)
