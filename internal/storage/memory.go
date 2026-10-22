package storage

import (
	"errors"
	"sync"
	"taskboard/internal/task"
)

// ErrNotFound — задача не найдена. Отдельная ошибка, чтобы хендлер мог отдать 404.
var ErrNotFound = errors.New("task not found")

// MemoryStorage хранит задачи в map. Потокобезопасен.
type MemoryStorage struct {
	mu    sync.RWMutex
	tasks map[string]task.Task
}

// NewMemory создаёт пустое хранилище.
func NewMemory() *MemoryStorage {
	// TODO: инициализировать map
	return nil
}

// Create сохраняет задачу.
// TODO: взять Lock, положить в map.
func (s *MemoryStorage) Create(t task.Task) error { return nil }

// Get возвращает задачу по ID или ErrNotFound.
// TODO: взять RLock, достать из map.
func (s *MemoryStorage) Get(id string) (task.Task, error) { return task.Task{}, ErrNotFound }

// List возвращает все задачи (копию слайса!).
// TODO: RLock, собрать слайс из map.
func (s *MemoryStorage) List() []task.Task { return nil }

// Update заменяет задачу по ID.
// TODO: Lock, проверить наличие, заменить.
func (s *MemoryStorage) Update(t task.Task) error { return ErrNotFound }

// Delete удаляет задачу по ID.
// TODO: Lock, проверить наличие, удалить.
func (s *MemoryStorage) Delete(id string) error { return ErrNotFound }
