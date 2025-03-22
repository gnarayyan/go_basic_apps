// storage/memory.go
package storage

import "github.com/gnarayyan/go_basic_apps/task-management/task"

type MemoryStorage struct {
    tasks  map[int]*task.Task
    nextID int
}

func NewMemoryStorage() *MemoryStorage {
    return &MemoryStorage{tasks: make(map[int]*task.Task), nextID: 1}
}

func (s *MemoryStorage) AddTask(t *task.Task) {
    t.ID = s.nextID
    s.tasks[t.ID] = t
    s.nextID++
}

func (s *MemoryStorage) GetTask(id int) (*task.Task, bool) {
    t, exists := s.tasks[id]
    return t, exists
}

func (s *MemoryStorage) DeleteTask(id int) bool {
    delete(s.tasks, id)
	return true;
}

func (s *MemoryStorage) ListTasks() []*task.Task {
    var tasks []*task.Task
    for _, t := range s.tasks {
        tasks = append(tasks, t)
    }
    return tasks
}