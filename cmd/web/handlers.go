package main

import (
	"encoding/json"
	"net/http"
)

type CreateUserForm struct {
	DisplayName string `json:"displayName"`
}

func setDefaultHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func home(w http.ResponseWriter, r *http.Request) {
	setDefaultHeaders(w)
	w.Write([]byte(`{}`))
}

func createRoom(w http.ResponseWriter, r *http.Request) {
	setDefaultHeaders(w)
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method "+r.Method+" not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte(`{"id": 1}`))
}

func createuser(w http.ResponseWriter, r *http.Request) {
	var u CreateUserForm

	setDefaultHeaders(w)
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method "+r.Method+" not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.Body == nil {
		http.Error(w, `Param "displayName" must be set.`, http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if u.DisplayName == "" {
		http.Error(w, `Param "displayName" must be set.`, http.StatusBadRequest)
		return
	}
	w.Write([]byte(`{"id": 1}`))
}
