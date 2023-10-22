package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)


func TestAddTaskHandler(t *testing.T) {
    queue = TaskQueue{tasks: make(map[int]Task)}

    r := mux.NewRouter()

    r.HandleFunc("/addTask", AddTaskHandler).Methods("POST")

    testServer := httptest.NewServer(r)
    defer testServer.Close()

    url := testServer.URL + "/addTask?n=8&d=2.0&n1=5.0&I=2.5&TTL=5"
    req, err := http.NewRequest("POST", url, nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()

    AddTaskHandler(rr, req)

    if rr.Code != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
    }
}