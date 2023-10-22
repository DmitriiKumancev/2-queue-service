package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

func main() {
	maxParallelTasks := flag.Int("maxParallelTasks", 5, "Максимальное количество параллельных задач")
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/addTask", AddTaskHandler).Methods("POST")
	r.HandleFunc("/getTasks", GetTasksHandler).Methods("GET")

	taskChannel := make(chan Task, *maxParallelTasks)

	queue = TaskQueue{tasks: make(map[int]Task)}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Printf("Server is running. Максимальное количество параллельных задач: %d\n", *maxParallelTasks)
		http.Handle("/", r)  
		if err := http.ListenAndServe(":8085", r); err != nil {
			fmt.Printf("Failed to start the server: %v\n", err)
			return
		}
	}()

	go ExecuteTasks(taskChannel)

	wg.Wait()
}




