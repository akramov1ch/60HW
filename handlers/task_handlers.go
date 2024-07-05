package handlers

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/Masterminds/squirrel"
    "net/http"
    "strconv"
    "60HW/db"
    "60HW/models"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    query := squirrel.Select("*").From("tasks")
    sql, args, err := query.ToSql()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    rows, err := db.DB.Query(sql, args...)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    tasks := []models.Task{}
    for rows.Next() {
        var task models.Task
        if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Done); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        tasks = append(tasks, task)
    }
    json.NewEncoder(w).Encode(tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    query := squirrel.Select("*").From("tasks").Where(squirrel.Eq{"id": id})
    sql, args, err := query.ToSql()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var task models.Task
    err = db.DB.QueryRow(sql, args...).Scan(&task.ID, &task.Title, &task.Description, &task.Done)
    if err != nil {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(task)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var task models.Task
    json.NewDecoder(r.Body).Decode(&task)

    query := squirrel.Insert("tasks").Columns("title", "description", "done").
        Values(task.Title, task.Description, task.Done).
        Suffix("RETURNING id")

    sql, args, err := query.ToSql()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = db.DB.QueryRow(sql, args...).Scan(&task.ID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    var task models.Task
    json.NewDecoder(r.Body).Decode(&task)

    query := squirrel.Update("tasks").
        Set("title", task.Title).
        Set("description", task.Description).
        Set("done", task.Done).
        Where(squirrel.Eq{"id": id})

    sql, args, err := query.ToSql()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    _, err = db.DB.Exec(sql, args...)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    query := squirrel.Delete("tasks").Where(squirrel.Eq{"id": id})
    sql, args, err := query.ToSql()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    _, err = db.DB.Exec(sql, args...)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted"})
}
