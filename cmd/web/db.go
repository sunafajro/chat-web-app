package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO реализовать загрузку параметров из .env
const (
	host     = "192.168.2.199"
	port     = 5432
	user     = "app_user"
	password = "app_secret_password"
	dbname   = "app_db"
)

var db *gorm.DB

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

func (j *JSONB) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}

// колонка deleted_at добавлена так как при запросе db.First() gorm автоматом добавляет проверку на not null по ней
type Room struct {
	gorm.Model
	Id         int
	Settings   JSONB `gorm:"type:jsonb"`
	Status     string
	Created_at string
	Updated_at string
	Deleted_at string
}

func connectDb() bool {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	handleError(err)
	db = connection
	return true
}

func disconnectDb() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()
}

func checkDbConnection() {
	if connectDb() {
		log.Println("Подключение к DB успешно.")
	}
}

func createRoom() Room {
	now := time.Now()
	room := Room{
		Settings:   JSONB{"property": "Value"},
		Status:     "active",
		Created_at: now.Format("2006-01-02 15:04:05"),
		Updated_at: now.Format("2006-01-02 15:04:05"),
	}
	// колонка deleted_at исключена иначе туда будет записываться пустая строка и это вызовет ошибку
	result := db.Table("chat_rooms").Omit("Deleted_at").Create(&room)
	handleError(result.Error)

	return room
}

func selectRoomById(id int) Room {
	var room Room
	db.Table("chat_rooms").First(&room, id)

	return room
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
