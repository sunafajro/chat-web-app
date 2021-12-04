package main

import (
	"net/http"
)

func getRoomHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method "+r.Method+" not allowed", http.StatusMethodNotAllowed)
		return
	}
	var roomId = r.URL.Query().Get("id")
	if roomId == "" {
		http.Error(w, `Param "id" must be set and not empty.`, http.StatusBadRequest)
		return
	}
	w.Write([]byte(`{"id": 1}`))
}
