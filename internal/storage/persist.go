package storage

// SaveToFile сохраняет все задачи в JSON-файл.
// TODO: взять s.List(), json.MarshalIndent, os.WriteFile(path, data, 0644).
func SaveToFile(s *MemoryStorage, path string) error { return nil }

// LoadFromFile загружает задачи из JSON-файла в хранилище.
// Если файла нет — не ошибка, просто пусто.
// TODO: os.ReadFile, если os.IsNotExist — вернуть nil; json.Unmarshal; пройтись и s.Create.
func LoadFromFile(s *MemoryStorage, path string) error { return nil }