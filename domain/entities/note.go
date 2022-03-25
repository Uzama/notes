package entities

import "time"

type Note struct {
	ID          int64
	UserID      int64
	Title       string
	Description string
	Archived    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
