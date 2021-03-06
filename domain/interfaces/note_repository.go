package interfaces

import (
	"context"
	"notes/domain/entities"
)

type NoteRepository interface {
	GetAllUnarchivedNotes(ctx context.Context, paginator entities.Paginator) ([]entities.Note, error)
	GetAllArchviedNotes(ctx context.Context, paginator entities.Paginator) ([]entities.Note, error)
	GetSingle(ctx context.Context, id int64) (entities.Note, error)
	CreateNote(ctx context.Context, note entities.Note) (int64, error)
	ArchiveNote(ctx context.Context, id int64) (bool, error)
	UnArchiveNote(ctx context.Context, id int64) (bool, error)
	UpdateNote(ctx context.Context, id int64, newNote entities.Note) (bool, error)
	DeleteNote(ctx context.Context, id int64) (bool, error)
	IsExists(ctx context.Context, id int64) (bool, error)
	IsArchived(ctx context.Context, id int64) (bool, error)
	RefreshCache()
}
