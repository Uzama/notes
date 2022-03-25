package router

import (
	"net/http"
	"notes/http/controllers"
	"notes/utils/container"

	"github.com/gorilla/mux"
)

func Init(ctr container.Containers) *mux.Router {

	r := mux.NewRouter()

	noteController := controllers.NewNoteController(ctr)

	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

	}).Methods(http.MethodGet)

	r.HandleFunc("note/get_un_archived", noteController.GetAllUnarchivedNotes).Methods(http.MethodGet)
	r.HandleFunc("note/get_archived", noteController.GetAllArchviedNotes).Methods(http.MethodGet)
	r.HandleFunc("note/create", noteController.CreateNote).Methods(http.MethodPost)
	r.HandleFunc("note/archive/{id}", noteController.ArchiveNote).Methods(http.MethodGet)
	r.HandleFunc("note/update/{id}", noteController.UpdateNote).Methods(http.MethodPost)
	r.HandleFunc("note/delete/{id}", noteController.DeleteNote).Methods(http.MethodGet)

	return r
}
