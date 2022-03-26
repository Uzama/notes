package controllers

import (
	"encoding/json"
	"net/http"
	"notes/domain/usecases"
	"notes/http/errors"
	"notes/http/transport/request"
	"notes/http/transport/request/decoders"
	"notes/http/transport/response"
	"notes/utils/container"
	"strconv"

	"github.com/gorilla/mux"
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

	ctx := r.Context()

	param := r.FormValue("paginator")

	pageDecoder := decoders.Paginator{}

	err := json.Unmarshal([]byte(param), &pageDecoder)
	if err != nil {
		errors.HandleError(w, err, http.StatusBadRequest)
		return
	}

	paginator, err := pageDecoder.Validate()
	if err != nil {
		errors.HandleError(w, err, http.StatusBadRequest)
		return
	}

	notes, err := ctl.usecase.GetAllUnarchivedNotes(ctx, paginator)
	if err != nil {
		errors.HandleError(w, err, http.StatusUnprocessableEntity)
		return
	}

	payload := response.Encode(notes, nil, "true")

	response.Send(w, payload, http.StatusOK)
}

func (ctl NoteController) GetAllArchviedNotes(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	param := r.FormValue("paginator")

	pageDecoder := decoders.Paginator{}

	err := json.Unmarshal([]byte(param), &pageDecoder)
	if err != nil {
		errors.HandleError(w, err, http.StatusBadRequest)
		return
	}

	paginator, err := pageDecoder.Validate()
	if err != nil {
		errors.HandleError(w, err, http.StatusBadRequest)
		return
	}

	notes, err := ctl.usecase.GetAllArchviedNotes(ctx, paginator)
	if err != nil {
		errors.HandleError(w, err, http.StatusUnprocessableEntity)
		return
	}

	payload := response.Encode(notes, nil, "true")

	response.Send(w, payload, http.StatusOK)
}

func (ctl NoteController) CreateNote(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	noteDecoder := decoders.Note{}

	err := request.Decode(ctx, r, &noteDecoder)
	if err != nil {
		errors.HandleError(w, err, http.StatusBadRequest)
		return
	}

	note, err := noteDecoder.Validate()
	if err != nil {
		errors.HandleError(w, err, http.StatusBadRequest)
		return
	}

	id, err := ctl.usecase.CreateNote(ctx, note)
	if err != nil {
		errors.HandleError(w, err, http.StatusUnprocessableEntity)
		return
	}

	payload := response.Encode(id, nil, "true")

	response.Send(w, payload, http.StatusCreated)
}

func (ctl NoteController) ArchiveNote(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errors.HandleError(w, err, http.StatusBadRequest)
		return
	}

	archived, err := ctl.usecase.ArchiveNote(ctx, int64(id))
	if err != nil {
		errors.HandleError(w, err, http.StatusUnprocessableEntity)
		return
	}

	payload := response.Encode(id, nil, archived)

	response.Send(w, payload, http.StatusAccepted)
}

func (ctl NoteController) UpdateNote(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errors.HandleError(w, err, http.StatusBadRequest)
		return
	}

	updateNoteDecoder := decoders.UpdateNote{}

	err = request.Decode(ctx, r, &updateNoteDecoder)
	if err != nil {
		errors.HandleError(w, err, http.StatusBadRequest)
		return
	}

	note, err := updateNoteDecoder.Validate()
	if err != nil {
		errors.HandleError(w, err, http.StatusBadRequest)
		return
	}

	updated, err := ctl.usecase.UpdateNote(ctx, int64(id), note)
	if err != nil {
		errors.HandleError(w, err, http.StatusUnprocessableEntity)
		return
	}

	payload := response.Encode(id, nil, updated)

	response.Send(w, payload, http.StatusAccepted)
}

func (ctl NoteController) DeleteNote(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errors.HandleError(w, err, http.StatusBadRequest)
		return
	}

	deleted, err := ctl.usecase.DeleteNote(ctx, int64(id))
	if err != nil {
		errors.HandleError(w, err, http.StatusUnprocessableEntity)
		return
	}

	payload := response.Encode(id, nil, deleted)

	response.Send(w, payload, http.StatusAccepted)
}
