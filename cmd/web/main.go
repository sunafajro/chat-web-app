package main

import (
	"log"
	"net/http"
)

func main() {
	checkDbConnection()

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/create-room", createRoomHandler)
	mux.HandleFunc("/create-user", createUserHandler)
	mux.HandleFunc("/sync-room-users", syncRoomUsersHandler)
	mux.HandleFunc("/send-message", sendMessageHandler)
	mux.HandleFunc("/get-room", getRoomHandler)
	mux.HandleFunc("/get-message", getMessageHandler)
	// mux.HandleFunc("/get-messages", getMessageHandler)
	// mux.HandleFunc("/get-room-messages", getMessageHandler)
	mux.HandleFunc("/read-message", readMessageHandler)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
