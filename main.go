package main

import (
	"log"
	"net/http"

	"com/fiuza/simple-go-mod/config"
	"com/fiuza/simple-go-mod/handlers"
	"com/fiuza/simple-go-mod/models"

	"github.com/gorilla/mux"
)

func main() {
	dbConnection := config.SetupDatabase()

	_, err := dbConnection.Exec(models.CreateTableSQL)

	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	taskHandler := handlers.NewTaskHandler(dbConnection)

	router.HandleFunc("/tasks", taskHandler.ReadTasks).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	defer dbConnection.Close()

	log.Fatal(http.ListenAndServe(":8080", router))
}
