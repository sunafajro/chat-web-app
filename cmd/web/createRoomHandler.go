package main

import "net/http"

func createRoomHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method "+r.Method+" not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte(`{"id": 1}`))
}
