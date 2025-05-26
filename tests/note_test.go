package tests

import (
	"bytes"
	"encoding/json"
	"go-notes-api/handlers"
	"go-notes-api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateNote(t *testing.T) {
	reqBody := map[string]string{
		"title":   "Test Başlık",
		"content": "Test İçerik",
	}
	jsonBody, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/notes", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/notes", handlers.CreateNote).Methods("POST")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Yanlis status kod beklenen %v gelen %v", http.StatusOK, status)
	}

	var noteResp map[string]string
	json.Unmarshal(rr.Body.Bytes(), &noteResp)

	if noteResp["title"] != "Test Başlık" {
		t.Errorf("Beklenen başlık 'Test Başlık', gelen %v", noteResp["title"])
	}

}

func TestGetAllNotes(t *testing.T) {

	note := map[string]string{"title": "Test Başlık", "content": "Test İçerik"}
	body, _ := json.Marshal(note)
	req, _ := http.NewRequest("POST", "/notes", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/notes", handlers.CreateNote).Methods("POST")
	router.ServeHTTP(rr, req)

	req2, _ := http.NewRequest("GET", "/notes", nil)
	rr2 := httptest.NewRecorder()
	router.HandleFunc("/notes", handlers.GetAllNotes).Methods("GET")
	router.ServeHTTP(rr2, req2)

	if rr2.Code != http.StatusOK {
		t.Errorf("Yanlış status kodu: beklenen %v, gelen %v", http.StatusOK, rr2.Code)
	}

	var notes []models.Note
	if err := json.Unmarshal(rr2.Body.Bytes(), &notes); err != nil {
		t.Fatalf("Response çözümlenemedi: %v", err)
	}

	if len(notes) == 0 {
		t.Errorf("Beklenen en az 1 not, ancak hiç yok.")
	}
}

func TestDeleteNote(t *testing.T) {

	note := map[string]string{"title": "Silinecek", "content": "İçerik"}
	body, _ := json.Marshal(note)
	req, _ := http.NewRequest("POST", "/notes", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/notes", handlers.CreateNote).Methods("POST")
	router.ServeHTTP(rr, req)

	var createdNote models.Note
	json.Unmarshal(rr.Body.Bytes(), &createdNote)

	deleteReq, _ := http.NewRequest("DELETE", "/notes/"+createdNote.ID, nil)
	deleteRR := httptest.NewRecorder()
	router.HandleFunc("/notes/{id}", handlers.DeleteNote).Methods("DELETE")
	router.ServeHTTP(deleteRR, deleteReq)

	if deleteRR.Code != http.StatusNoContent {
		t.Errorf("Beklenen 204 No Content, gelen %v", deleteRR.Code)
	}
}
