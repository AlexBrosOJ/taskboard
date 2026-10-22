package task

import (
	"errors"
	"time"
)

// Status — статус задачи. Строковый тип, чтобы JSON был читаемым.
type Status string

const (
	StatusTodo  Status = "todo"
	StatusDone  Status = "done"
)

// Task — одна задача.
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Status      Status    `json:"status"`
	Tags        []string  `json:"tags,omitempty"`
	Deadline    time.Time `json:"deadline,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Validate проверяет, что задача корректна.
// TODO: title не пустой; status либо todo, либо done; deadline не в прошлом (если задан).
func (t Task) Validate() error {
	// подсказка: errors.New / fmt.Errorf
	return nil
}

// New создаёт новую задачу с заполненными CreatedAt/UpdatedAt и статусом todo.
// TODO: сгенерировать ID (например, через crypto/rand или time.Now().UnixNano()).
func New(title, description string, tags []string, deadline time.Time) (Task, error) {
	// подсказка: собрать Task, вызвать Validate, вернуть ошибку если что
	return Task{}, errors.New("not implemented")
}