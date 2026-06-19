package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"manager/internal/model"
	"manager/internal/service"
)

type TaskHandler struct {
	service service.ItaskService
}

func New(service service.ItaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		log.Println("decode task:", err)
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)
		return
	}

	if err := h.service.Create(task); err != nil {
		log.Println("create task:", err)
		http.Error(
			w,
			"internal server error",
			http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.GetAll()
	if err != nil {
		log.Println("get all tasks:", err)
		http.Error(
			w,
			"internal server error",
			http.StatusInternalServerError,
		)
		return
	}

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		log.Println("encode tasks:", err)
		http.Error(
			w,
			"internal server error",
			http.StatusInternalServerError,
		)
		return
	}
}

func (h *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Println("invalid id:", err)
		http.Error(
			w,
			"invalid id",
			http.StatusBadRequest,
		)
		return
	}

	task, err := h.service.GetByID(int(id))
	if err != nil {
		log.Println("get task by id:", err)
		http.Error(
			w,
			"task not found",
			http.StatusNotFound,
		)
		return
	}

	if err := json.NewEncoder(w).Encode(task); err != nil {
		log.Println("encode task:", err)
		http.Error(
			w,
			"internal server error",
			http.StatusInternalServerError,
		)
		return
	}
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Println("invalid id:", err)
		http.Error(
			w,
			"invalid id",
			http.StatusBadRequest,
		)
		return
	}

	if err := h.service.Delete(int(id)); err != nil {
		log.Println("delete task:", err)
		http.Error(
			w,
			"internal server error",
			http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}