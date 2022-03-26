package decoders

import (
	"errors"
	"notes/domain/entities"
)

type UpdateNote struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (note UpdateNote) Format() string {
	return `
		{
			"title": "title",
			"description": "description"
		}
	`
}

func (note UpdateNote) Validate() (entities.Note, error) {

	n := entities.Note{}

	if note.Description == "" && note.Title == "" {
		return n, errors.New("either title or description must be given")
	}

	n.Title = note.Title
	n.Description = note.Description

	return n, nil
}
