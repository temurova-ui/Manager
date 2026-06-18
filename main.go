package main

import (
	"fmt"
	"log"
	"manager/internal/config"
	"manager/internal/handler"
	"manager/internal/model"
	"manager/internal/repository"
	"manager/internal/service"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.New("./config/config.env")
	if err != nil{
	  log.Fatal("config.New",err)
	}
	
	dsn := fmt.Sprintf(
	  "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
	  cfg.DBHost,
	  cfg.DBUser,
	  cfg.DBPassword,
	  cfg.DBName,
	  cfg.DBPort,
	)

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
	// mux.HandleFunc("PUT /tasks/{id}", taskHandler.Update)
	mux.HandleFunc("DELETE /tasks/{id}", taskHandler.Delete)

	log.Println("Server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", mux))	
}