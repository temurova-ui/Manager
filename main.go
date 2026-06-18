package main

import (
	"log"
	"net/http"

	"manager/internal/handler"
	"manager/internal/model"
	"manager/internal/repository"
	"manager/internal/service"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=1234 dbname=manager port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.Task{})

	repo := repository.New(db)
	srv := service.New(repo)
	taskHandler := handler.New(srv)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /tasks", taskHandler.Create)
	mux.HandleFunc("GET /tasks", taskHandler.GetAll)
	mux.HandleFunc("GET /tasks/{id}", taskHandler.GetByID)
	mux.HandleFunc("PUT /tasks/{id}", taskHandler.Update)
	mux.HandleFunc("DELETE /tasks/{id}", taskHandler.Delete)

	log.Println("Server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", mux))
}