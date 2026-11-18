package internal

type Status string

const (
	Todo       Status = "todo"
	InProgress Status = "in-progress"
	Done       Status = "done"
)

const JsonStoreName = "store.json"
