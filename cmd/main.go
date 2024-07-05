package main

import (
    "log"
    "net/http"
    "60HW/db"
    "60HW/handlers"

    "github.com/gorilla/mux"
)

func main() {
    db.InitDB("user=postgres password=vakhaboff dbname=shaxboz sslmode=disable")

    router := mux.NewRouter()

    router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
    router.HandleFunc("/tasks/{id}", handlers.GetTask).Methods("GET")
    router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
    router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
    router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8000", router))
}
