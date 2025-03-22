package task

import "time"

type Task struct {
    ID        int
    Desc      string
    Status    string
    CreatedAt time.Time
}

type Storage interface {
    AddTask(t *Task)
    GetTask(id int) (*Task, bool)
    DeleteTask(id int)
    ListTasks() []*Task
}