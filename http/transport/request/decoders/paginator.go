package decoders

import (
	"errors"
	"notes/domain/entities"
)

type Paginator struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

func (paginator Paginator) Format() string {
	return `
		{
			"page": 1,
			"size": 10
		}
	`
}

func (paginator Paginator) Validate() (entities.Paginator, error) {

	p := entities.Paginator{}

	if paginator.Page < 1 {
		return p, errors.New("page must be a positive integer")
	}

	if paginator.Size < 10 || paginator.Size > 100 {
		return p, errors.New("size must be inbetween 10 - 100")
	}

	p.Page = paginator.Page
	p.Size = paginator.Size

	return p, nil
}
