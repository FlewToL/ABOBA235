package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"example.com/pz6-gorm/internal/db"
	httpapi "example.com/pz6-gorm/internal/httpapi"
	"example.com/pz6-gorm/internal/models"
)

func main() {
	// Загружаем настройки из .env в корне проекта.
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found; using system environment")
	}

	d := db.Connect()

	// Создаём таблицы и необходимые связи.
	if err := d.AutoMigrate(
		&models.User{},
		&models.Note{},
		&models.Tag{},
	); err != nil {
		log.Fatal("migrate:", err)
	}

	r := httpapi.BuildRouter(d)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}