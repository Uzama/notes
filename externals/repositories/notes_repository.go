package repositories

import (
	"context"
	"database/sql"
	"notes/domain/entities"
	"notes/domain/interfaces"
)

type NoteRepository struct {
	db *sql.DB
}

func NewNoteRepository(db *sql.DB) interfaces.NoteRepository {
	repo := &NoteRepository{
		db: db,
	}

	return repo
}

func (repo NoteRepository) GetAllUnarchivedNotes(ctx context.Context, paginator entities.Paginator) ([]entities.Note, error) {
	return nil, nil
}

func (repo NoteRepository) GetAllArchviedNotes(ctx context.Context, paginator entities.Paginator) ([]entities.Note, error) {
	return nil, nil
}

func (repo NoteRepository) CreateNote(ctx context.Context, note entities.Note) (int64, error) {
	return 0, nil
}

func (repo NoteRepository) ArchiveNote(ctx context.Context, id int64) (bool, error) {
	return false, nil
}

func (repo NoteRepository) UpdateNote(ctx context.Context, id int64, newNote entities.Note) (bool, error) {
	return false, nil
}

func (repo NoteRepository) DeleteNote(ctx context.Context, id int64) (bool, error) {
	return false, nil
}
