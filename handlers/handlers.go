package handlers

import (
	"encoding/json"
	"go-notes-api/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var notes []models.Note

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(notes)
}

func GetNoteByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, note := range notes {
		if note.ID == params["id"] {
			json.NewEncoder(w).Encode(note)
			return
		}
	}
	http.Error(w, "Not bulunamadi", http.StatusNotFound)
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	_ = json.NewDecoder(r.Body).Decode(&note)
	note.ID = uuid.New().String()
	notes = append(notes, note)
	json.NewEncoder(w).Encode(note)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, note := range notes {
		if note.ID == params["id"] {
			notes = append(notes[:index], notes[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
		}
	}

	http.Error(w, "Silincek Not bulunmamaktadir", http.StatusNotFound)
}
