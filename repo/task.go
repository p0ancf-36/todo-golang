package repo

import "todo-golang.com/domain"

type Task interface {
	Get(id int) (*domain.Task, error)
	GetPage(page int) ([]domain.Task, error)
	Create(title, description string) (*domain.Task, error)
	Delete(id int) error
}
