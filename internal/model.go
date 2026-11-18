package internal

import (
	"time"
)

type Task struct {
	Id          int
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
