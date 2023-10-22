package main

import (
    "time"
)

var executingTasks = make(map[int]struct{})
var queue TaskQueue

func (q *TaskQueue) AddTask(task Task) {
	q.Lock()
	q.tasks[task.ID] = task
	q.Unlock()
}

func (q *TaskQueue) GetTasks() []Task {
	q.Lock()
	tasks := make([]Task, 0, len(q.tasks))
	for _, task := range q.tasks {
		tasks = append(tasks, task)
	}
	q.Unlock()
	return tasks
}

func ExecuteTasks(taskChannel chan Task) {
	for {
		queue.Lock()
		tasksToExecute := []Task{}
		tasksToDelete := []int{}
		for id, task := range queue.tasks {
			if task.Status == "В очереди" {
				tasksToExecute = append(tasksToExecute, task)
				if len(tasksToExecute) >= cap(taskChannel) {
					break
				}
			} else if task.Status == "Завершена" && task.EndTime.Add(time.Duration(task.TTL)*time.Second).Before(time.Now()) {
				tasksToDelete = append(tasksToDelete, id)
			}
		}
		queue.Unlock()

		for _, task := range tasksToExecute {
			queue.Lock()
			queue.tasks[task.ID] = Task{
				ID:           task.ID,
				Status:       "В процессе",
				N:            task.N,
				D:            task.D,
				N1:           task.N1,
				I:            task.I,
				TTL:          task.TTL,
				Iterations:   task.Iterations,
				StartTime:    time.Now(),
				EndTime:      task.EndTime,
				CurrentValue: task.CurrentValue,
			}
			queue.Unlock()

			go func(task Task) {
				ticker := time.NewTicker(time.Duration(int(task.I*1000)) * time.Millisecond)
				defer ticker.Stop()

				for range ticker.C {
					if task.Iterations < task.N {
						task.CurrentValue += task.D
						task.Iterations++
					} else {
						break
					}
				}

				queue.Lock()
				queue.tasks[task.ID] = Task{
					ID:           task.ID,
					Status:       "Завершена",
					N:            task.N,
					D:            task.D,
					N1:           task.N1,
					I:            task.I,
					TTL:          task.TTL,
					Iterations:   task.Iterations,
					StartTime:    task.StartTime,
					EndTime:      time.Now(),
					CurrentValue: task.CurrentValue,
				}
				queue.Unlock()

				<-taskChannel
				delete(executingTasks, task.ID)
			}(task)
			taskChannel <- task
		}

		queue.Lock()
		for _, id := range tasksToDelete {
			delete(queue.tasks, id)
		}
		tasksToDelete = nil
		queue.Unlock()
	}
}
