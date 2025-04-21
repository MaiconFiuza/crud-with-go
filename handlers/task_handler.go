package handlers

import (
	"com/fiuza/simple-go-mod/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TaskHandler struct {
	DB *sql.DB
}

func NewTaskHandler(db *sql.DB) *TaskHandler {
	return &TaskHandler{DB: db}
}

func (TaskHandler *TaskHandler) ReadTasks(writer http.ResponseWriter, request *http.Request) {
	rows, err := TaskHandler.DB.Query("SELECT * FROM tasks")

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var tasks = make([]models.Task, 0)

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		tasks = append(tasks, task)
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tasks)
}

func (TaskHandler *TaskHandler) CreateTask(writer http.ResponseWriter, request *http.Request) {
	var task models.Task
	err := json.NewDecoder(request.Body).Decode(&task)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id"
	err = TaskHandler.DB.QueryRow(query, task.Title, task.Description, task.Status).Scan(&task.ID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(task)
}

func (TaskHandler *TaskHandler) UpdateTask(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, "Invalid Task ID", http.StatusBadRequest)
		return
	}

	var task models.Task
	err = json.NewDecoder(request.Body).Decode(&task)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := TaskHandler.DB.Exec("UPDATE tasks SET title = $1, description = $2, status = $3 WHERE id = $4", task.Title, task.Description, task.Status, id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(writer, "No Task Found with this ID", http.StatusNotFound)
	}

	task.ID = id
	writer.Header().Set("Content-type", "application/json")
}

func (TaskHandler *TaskHandler) DeleteTask(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, "Invalid Task ID", http.StatusBadRequest)
		return
	}

	result, err := TaskHandler.DB.Exec("DELETE FROM tasks WHERE id = $1", id)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(writer, "No Task Found with this ID", http.StatusNotFound)
	}

	writer.WriteHeader(http.StatusNoContent)
}
