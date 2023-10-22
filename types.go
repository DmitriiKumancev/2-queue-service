package main

import (
    "sync"
    "time"
)

type Task struct {
	ID           int
	Status       string
	N            int
	D            float64
	N1           float64
	I            float64
	TTL          int
	Iterations   int
	StartTime    time.Time
	EndTime      time.Time
	CurrentValue float64
}

type TaskQueue struct {
	sync.Mutex
	tasks map[int]Task
}



