package main

import (
	"encoding/json"
	"net/http"
)

type SyncRoomUsersForm struct {
	UserIds []int `json:"userIds"`
}

func syncRoomUsersHandler(w http.ResponseWriter, r *http.Request) {
	var data SyncRoomUsersForm

	w.Header().Set("Content-Type", "application/json")
	var roomId = r.URL.Query().Get("id")
	if roomId == "" {
		http.Error(w, `Param "id" must be set and not empty.`, http.StatusBadRequest)
		return
	}
	if r.Body == nil {
		http.Error(w, `Param "userIds" must be set and not empty.`, http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(data.UserIds) == 0 {
		http.Error(w, `Param "userIds" must be set and not empty.`, http.StatusBadRequest)
		return
	}

	w.Write([]byte(`{"result": true}`))
}
