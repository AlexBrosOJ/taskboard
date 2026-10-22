package main

import (
	"log"
	"net/http"

	"taskboard/internal/httpapi"
	"taskboard/internal/service"
	"taskboard/internal/storage"
)

func main() {
	const dataFile = "tasks.json"
	const addr = ":8080"

	store := storage.NewMemory()
	// TODO: storage.LoadFromFile(store, dataFile) — и залогировать ошибку

	svc := service.New(store)
	api := httpapi.New(svc)

	mux := http.NewServeMux()
	api.Register(mux)

	log.Printf("listening on %s", addr)
	// TODO: http.ListenAndServe(addr, mux)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
