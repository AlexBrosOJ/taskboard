package httpapi

import (
	"net/http"

	"taskboard/internal/service"
)

// API держит зависимости хендлеров.
type API struct {
	svc *service.Service
}

func New(svc *service.Service) *API { return &API{svc: svc} }

// Register вешает хендлеры на mux.
// TODO: mux.HandleFunc("POST /tasks", a.create) и т.д. (Go 1.22+ умеет методы в паттерне)
func (a *API) Register(mux *http.ServeMux) {}

// create — POST /tasks. Принимает JSON, возвращает 201 + задачу.
// TODO: json.Decode, вызвать svc.Add, json.Encode, WriteHeader(201).
func (a *API) create(w http.ResponseWriter, r *http.Request) {}

// list — GET /tasks?status=done&tag=work.
// TODO: прочитать query, вызвать svc.List, отдать JSON.
func (a *API) list(w http.ResponseWriter, r *http.Request) {}

// get — GET /tasks/{id}. 404 если нет.
// TODO: r.PathValue("id"), svc... (нужен метод Get в сервисе — добавь)
func (a *API) get(w http.ResponseWriter, r *http.Request) {}

// markDone — PATCH /tasks/{id}/done.
func (a *API) markDone(w http.ResponseWriter, r *http.Request) {}

// remove — DELETE /tasks/{id}. 204 при успехе.
func (a *API) remove(w http.ResponseWriter, r *http.Request) {}

// writeJSON — маленький хелпер: ставит Content-Type, кодирует, пишет статус.
// TODO: реализовать.
func writeJSON(w http.ResponseWriter, status int, v any) {}
