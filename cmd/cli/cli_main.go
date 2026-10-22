package main

import (
	"flag"
	"fmt"
)

func main() {
	// Пример парсинга флагов.
	// Запуск: go run ./cmd/cli -title "Купить хлеб"
	title := flag.String("title", "", "Заголовок новой задачи")
	flag.Parse()

	if *title == "" {
		fmt.Println("Привет! Я CLI-клиент TaskBoard.")
		fmt.Println("Пока я ничего не умею. Твоя задача — научить меня делать HTTP-запросы к серверу.")
		return
	}

	fmt.Printf("Ты хочешь добавить задачу: %s\n", *title)
	fmt.Println("Осталось написать отправку этого запроса на сервер!")
}
