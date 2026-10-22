package task

import "testing"

func TestValidate(t *testing.T) {
	// Шпаргалка: так пишутся табличные тесты.
	// TODO: добавь свои кейсы (пустой title, статус done и т.д.)
	tests := []struct {
		name    string
		task    Task
		wantErr bool
	}{
		// {"пустой тайтл", Task{Title: ""}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.task.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}