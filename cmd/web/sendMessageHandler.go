package main

import (
	"encoding/json"
	"net/http"
)

type CreateMessageForm struct {
	UserId    int    `json:"userId"`
	RoomId    int    `json:"roomId"`
	Body      string `json:"body"`
	RepliedTo int    `json:"repliedTo"`
}

func sendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var data CreateMessageForm

	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method "+r.Method+" not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.Body == nil {
		http.Error(w, `Params "userId", "roomId", "body" must be set and not empty.`, http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if data.UserId == 0 {
		http.Error(w, `Param "userId" must be set and not empty.`, http.StatusBadRequest)
		return
	}
	if data.RoomId == 0 {
		http.Error(w, `Param "roomId" must be set and not empty.`, http.StatusBadRequest)
		return
	}
	if data.Body == "" {
		http.Error(w, `Param "body" must be set and not empty.`, http.StatusBadRequest)
		return
	}
	w.Write([]byte(`{"id": 1}`))
}
