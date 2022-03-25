package controllers

import (
	"net/http"
	"notes/domain/usecases"
	"notes/utils/container"
)

type NoteController struct {
	usecase usecases.NoteUsecase
}

func NewNoteController(ctr container.Containers) NoteController {
	ctl := NoteController{
		usecase: usecases.NewNoteUsecase(ctr.Repositories.Note),
	}

	return ctl
}

func (ctl NoteController) GetAllUnarchivedNotes(w http.ResponseWriter, r *http.Request) {

}

func (ctl NoteController) GetAllArchviedNotes(w http.ResponseWriter, r *http.Request) {

}

func (ctl NoteController) CreateNote(w http.ResponseWriter, r *http.Request) {

}

func (ctl NoteController) ArchiveNote(w http.ResponseWriter, r *http.Request) {

}

func (ctl NoteController) UpdateNote(w http.ResponseWriter, r *http.Request) {

}

func (ctl NoteController) DeleteNote(w http.ResponseWriter, r *http.Request) {

}
