package domain

// Status enum type for task status
type Status int8

const (
	StatusTodo Status = iota
	StatusDone
	StatusInProgress
)

// Priority enum type for task priority
type Priority int8

const (
	Pri4 Priority = iota
	Pri2
	Pri3
	Pri1
)
