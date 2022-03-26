package decoders

import (
	"errors"
	"notes/domain/entities"
)

type Note struct {
	UserID      int64  `json:"user_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (note Note) Format() string {
	return `
		{
			"user_id": 123,
			"title": "title",
			"description": "description"
		}
	`
}

func (note Note) Validate() (entities.Note, error) {

	n := entities.Note{}

	if note.UserID <= 0 {
		return n, errors.New("user_id must be a positive integer")
	}

	n.UserID = note.UserID
	n.Title = note.Title
	n.Description = note.Description

	return n, nil
}
