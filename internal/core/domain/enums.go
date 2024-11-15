package domain

// Status enum type for task status
type Status int8

const (
	_ Status = iota
	StatusTodo
	StatusDone
	StatusInProgress
)

// Priority enum type for task priority
type Priority int8

const (
	_ Priority = iota
	Pri4
	Pri2
	Pri3
	Pri1
)
