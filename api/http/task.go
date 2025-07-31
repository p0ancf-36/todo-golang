package http

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"todo-golang.com/domain"
	"todo-golang.com/repo"
	"todo-golang.com/views"
)

type Task struct {
	repo repo.Task
}

func NewTask(repo repo.Task) Task {
	return Task{
		repo: repo,
	}
}

func (h *Task) UseHandlers(r chi.Router) {
	r.Get("/", h.get)
}

func (h *Task) get(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	if query.Has("page") {
		if page, err := strconv.Atoi(query.Get("page")); err != nil {
			http.Error(w, "page is not a number", http.StatusBadRequest)
		} else {
			h.getTaskPage(w, r, page)
		}
	} else if query.Has("id") {
		h.getTaskById(w, r)
	}
}

func (h *Task) getTaskPage(w http.ResponseWriter, r *http.Request, page int) {
	type Tasks struct {
		Tasks []domain.Task
	}

	tasks, err := h.repo.GetPage(page)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	views.Render(w, "tasks", Tasks{Tasks: tasks})

	// if err = json.NewEncoder(w).Encode(tasks); err != nil {
	// 	http.Error(w, err.Error(), 500)
	// }
}

func (h *Task) getTaskById(w http.ResponseWriter, r *http.Request) {

}
