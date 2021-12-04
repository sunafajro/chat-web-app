package main

import (
	"encoding/json"
	"net/http"
)

type ReadMessageForm struct {
	UserId int `json:"userId"`
}

func readMessageHandler(w http.ResponseWriter, r *http.Request) {
	var data ReadMessageForm

	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method "+r.Method+" not allowed", http.StatusMethodNotAllowed)
		return
	}
	var messageId = r.URL.Query().Get("id")
	if messageId == "" {
		http.Error(w, `Param "id" must be set and not empty.`, http.StatusBadRequest)
		return
	}
	if r.Body == nil {
		http.Error(w, `Param "userId" must be set and not empty.`, http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(`{"id": 1}`))
}
