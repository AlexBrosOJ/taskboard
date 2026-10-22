package service

import (
	"taskboard/internal/storage"
	"taskboard/internal/task"
	"time"
)

// Service — бизнес-логика поверх хранилища.
type Service struct {
	store *storage.MemoryStorage
}

func New(store *storage.MemoryStorage) *Service {
	return &Service{store: store}
}

// Add валидирует и создаёт задачу.
// TODO: вызвать task.New, потом store.Create.
func (s *Service) Add(title, description string, tags []string, deadline time.Time) (task.Task, error) {
	return task.Task{}, nil
}

// List возвращает задачи, отфильтрованные по статусу и тегу (пустые = без фильтра).
// TODO: взять store.List(), пройтись циклом, отфильтровать.
func (s *Service) List(status task.Status, tag string) []task.Task { return nil }

// MarkDone переводит задачу в done и обновляет UpdatedAt.
// TODO: store.Get, поменять статус, store.Update.
func (s *Service) MarkDone(id string) (task.Task, error) { return task.Task{}, nil }

// Remove удаляет задачу.
func (s *Service) Remove(id string) error { return nil }
