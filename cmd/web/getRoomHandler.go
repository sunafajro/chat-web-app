package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func getRoomHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method "+r.Method+" not allowed", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if id == 0 || err != nil {
		http.Error(w, `Param "id" must be set and not empty.`, http.StatusBadRequest)
		return
	}

	var room = selectRoomById(id)

	settings, err := room.Settings.Value()
	var response = fmt.Sprintf(
		`{"id": %d, "settings": %s, "status": "%s", "created_at": "%s", "updated_at": "%s", "deleted_at": "%s"}`,
		room.Id, settings, room.Status, room.Created_at, room.Updated_at, room.Deleted_at)
	w.Write([]byte(response))
}
