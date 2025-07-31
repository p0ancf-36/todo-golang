package domain

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
}

func NewTask(id int, title, description string, isDone bool) Task {
	return Task{
		Id:          id,
		Title:       title,
		Description: description,
		IsDone:      isDone,
	}
}
