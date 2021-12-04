package main

import (
	"encoding/json"
	"net/http"
)

type CreateUserForm struct {
	DisplayName string `json:"displayName"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var data CreateUserForm

	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method "+r.Method+" not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.Body == nil {
		http.Error(w, `Param "displayName" must be set and not empty.`, http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if data.DisplayName == "" {
		http.Error(w, `Param "displayName" must be set and not empty.`, http.StatusBadRequest)
		return
	}
	w.Write([]byte(`{"id": 1}`))
}
