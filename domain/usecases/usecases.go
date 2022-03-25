package usecases

import (
	"context"
	"notes/domain/entities"
	"notes/domain/interfaces"
)

type NoteUsecase struct {
	repo interfaces.NoteRepository
}

func NewNoteUsecase(repo interfaces.NoteRepository) NoteUsecase {
	usecase := NoteUsecase{
		repo: repo,
	}

	return usecase
}

func (usecase NoteUsecase) GetAllUnarchivedNotes(ctx context.Context, paginator entities.Paginator) ([]entities.Note, error) {
	return nil, nil
}

func (usecase NoteUsecase) GetAllArchviedNotes(ctx context.Context, paginator entities.Paginator) ([]entities.Note, error) {
	return nil, nil
}

func (usecase NoteUsecase) CreateNote(ctx context.Context, note entities.Note) (int64, error) {
	return 0, nil
}

func (usecase NoteUsecase) ArchiveNote(ctx context.Context, id int64) (bool, error) {
	return false, nil
}

func (usecase NoteUsecase) UpdateNote(ctx context.Context, id int64, newNote entities.Note) (bool, error) {
	return false, nil
}

func (usecase NoteUsecase) DeleteNote(ctx context.Context, id int64) (bool, error) {
	return false, nil
}
