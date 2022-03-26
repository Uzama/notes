package usecases

import (
	"context"
	"errors"
	"fmt"
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

	return usecase.repo.GetAllUnarchivedNotes(ctx, paginator)
}

func (usecase NoteUsecase) GetAllArchviedNotes(ctx context.Context, paginator entities.Paginator) ([]entities.Note, error) {

	return usecase.repo.GetAllArchviedNotes(ctx, paginator)
}

func (usecase NoteUsecase) CreateNote(ctx context.Context, note entities.Note) (int64, error) {

	// validate max length
	if len(note.Title) > 120 {
		return 0, errors.New("title lenght is too long")
	}

	if len(note.Description) > 255 {
		return 0, errors.New("description lenght is too long")
	}

	// save the note
	id, err := usecase.repo.CreateNote(ctx, note)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (usecase NoteUsecase) ArchiveNote(ctx context.Context, id int64) (bool, error) {

	// check for existence
	isExists, err := usecase.repo.IsExists(ctx, id)
	if err != nil {
		return false, err
	}

	if !isExists {
		return false, fmt.Errorf("%d note is not exists", id)
	}

	// check whether note is already archived
	isArchived, err := usecase.repo.IsArchived(ctx, id)
	if err != nil {
		return false, err
	}

	if isArchived {
		return false, fmt.Errorf("%d note is already archived", id)
	}

	// archive the note
	archived, err := usecase.repo.ArchiveNote(ctx, id)
	if err != nil {
		return false, err
	}

	return archived, nil
}

func (usecase NoteUsecase) UnArchiveNote(ctx context.Context, id int64) (bool, error) {

	// check for existence
	isExists, err := usecase.repo.IsExists(ctx, id)
	if err != nil {
		return false, err
	}

	if !isExists {
		return false, fmt.Errorf("%d note is not exists", id)
	}

	// check whether note is already archived
	isArchived, err := usecase.repo.IsArchived(ctx, id)
	if err != nil {
		return false, err
	}

	if !isArchived {
		return false, fmt.Errorf("%d note is already un archived", id)
	}

	// archive the note
	unArchived, err := usecase.repo.UnArchiveNote(ctx, id)
	if err != nil {
		return false, err
	}

	return unArchived, nil
}

func (usecase NoteUsecase) UpdateNote(ctx context.Context, id int64, newNote entities.Note) (bool, error) {

	// validate max length
	if len(newNote.Title) > 120 {
		return false, errors.New("title lenght is too long")
	}

	if len(newNote.Description) > 255 {
		return false, errors.New("description lenght is too long")
	}

	// check for existence
	isExists, err := usecase.repo.IsExists(ctx, id)
	if err != nil {
		return false, err
	}

	if !isExists {
		return false, fmt.Errorf("%d note is not exists", id)
	}

	// get the old note
	oldNote, err := usecase.repo.GetSingle(ctx, id)
	if err != nil {
		return false, err
	}

	// update old values if new values note provided
	if newNote.Title == "" {
		newNote.Title = oldNote.Title
	}

	if newNote.Description == "" {
		newNote.Description = oldNote.Description
	}

	// update the note
	updated, err := usecase.repo.UpdateNote(ctx, id, newNote)
	if err != nil {
		return false, err
	}

	return updated, nil
}

func (usecase NoteUsecase) DeleteNote(ctx context.Context, id int64) (bool, error) {

	// check for existence
	isExists, err := usecase.repo.IsExists(ctx, id)
	if err != nil {
		return false, err
	}

	if !isExists {
		return false, fmt.Errorf("%d note is not exists", id)
	}

	// delete the note
	deleted, err := usecase.repo.DeleteNote(ctx, id)
	if err != nil {
		return false, err
	}

	return deleted, nil
}
