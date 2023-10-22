package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
    n, err := strconv.Atoi(r.FormValue("n"))
    if err != nil {
        http.Error(w, "Invalid 'n' value", http.StatusBadRequest)
        return
    }

    d, err := strconv.ParseFloat(r.FormValue("d"), 64)
    if err != nil {
        http.Error(w, "Invalid 'd' value", http.StatusBadRequest)
        return
    }

    n1, err := strconv.ParseFloat(r.FormValue("n1"), 64)
    if err != nil {
        http.Error(w, "Invalid 'n1' value", http.StatusBadRequest)
        return
    }

    I, err := strconv.ParseFloat(r.FormValue("I"), 64)
    if err != nil {
        http.Error(w, "Invalid 'I' value", http.StatusBadRequest)
        return
    }

    TTL, err := strconv.Atoi(r.FormValue("TTL"))
    if err != nil {
        http.Error(w, "Invalid 'TTL' value", http.StatusBadRequest)
        return
    }

    task := Task{
        ID:           len(queue.tasks) + 1,
        Status:       "В очереди",
        N:            n,
        D:            d,
        N1:           n1,
        I:            I,
        TTL:          TTL,
        Iterations:   0,
        StartTime:    time.Time{},
        EndTime:      time.Time{},
        CurrentValue: n1,
    }
    queue.AddTask(task)

    fmt.Printf("Добавлена задача %d в очередь.\n", task.ID)
    w.WriteHeader(http.StatusOK)
}


func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := queue.GetTasks()
    taskStatuses := make([]Task, len(tasks))

    copy(taskStatuses, tasks)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(taskStatuses)
}

