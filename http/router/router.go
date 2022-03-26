package router

import (
	"net/http"
	"notes/http/controllers"
	"notes/http/transport/response"
	"notes/utils/container"

	"github.com/gorilla/mux"
)

func Init(ctr container.Containers) *mux.Router {

	r := mux.NewRouter()

	noteController := controllers.NewNoteController(ctr)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response.Send(w, []byte("Notes APP"), http.StatusOK)
	}).Methods(http.MethodGet)

	r.HandleFunc("/note/get_un_archived", noteController.GetAllUnarchivedNotes).Methods(http.MethodGet)
	r.HandleFunc("/note/get_archived", noteController.GetAllArchviedNotes).Methods(http.MethodGet)
	r.HandleFunc("/note/create", noteController.CreateNote).Methods(http.MethodPost)
	r.HandleFunc("/note/archive/{id}", noteController.ArchiveNote).Methods(http.MethodGet)
	r.HandleFunc("/note/update/{id}", noteController.UpdateNote).Methods(http.MethodPut)
	r.HandleFunc("/note/delete/{id}", noteController.DeleteNote).Methods(http.MethodDelete)

	r.HandleFunc("/note/refresh_cache", func(w http.ResponseWriter, r *http.Request) {
		ctr.Repositories.Note.RefreshCache()
		response.Send(w, []byte("cache refreshed"), http.StatusOK)
	})

	return r
}
