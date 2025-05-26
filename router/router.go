package router

import (
	"go-notes-api/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/notes", handlers.GetAllNotes).Methods("GET")
	r.HandleFunc("/notes/{id}", handlers.GetNoteByID).Methods("GET")
	r.HandleFunc("/notes", handlers.CreateNote).Methods("POST")
	r.HandleFunc("/notes/{id}", handlers.DeleteNote).Methods("DELETE")
}
