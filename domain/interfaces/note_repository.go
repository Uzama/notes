package interfaces

import (
	"context"
	"notes/domain/entities"
)

type NoteRepository interface {
	GetAllUnarchivedNotes(ctx context.Context, paginator entities.Paginator) ([]entities.Note, error)
	GetAllArchviedNotes(ctx context.Context, paginator entities.Paginator) ([]entities.Note, error)
	CreateNote(ctx context.Context, note entities.Note) (int64, error)
	ArchiveNote(ctx context.Context, id int64) (bool, error)
	UpdateNote(ctx context.Context, id int64, newNote entities.Note) (bool, error)
	DeleteNote(ctx context.Context, id int64) (bool, error)
}
