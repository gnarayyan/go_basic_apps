// task/filter.go
package task

func Filter[T comparable](tasks []*Task, predicate func(*Task) T, value T) []*Task {
    var result []*Task
    for _, t := range tasks {
        if predicate(t) == value {
            result = append(result, t)
        }
    }
    return result
}

// Example usage in main.go:
// completedTasks := task.Filter(store.ListTasks(), func(t *task.Task) string { return t.Status }, "completed")