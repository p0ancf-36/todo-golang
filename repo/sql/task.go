package sql

import (
	"database/sql"

	"todo-golang.com/domain"
	"todo-golang.com/repo"
)

type Task struct {
	db *sql.DB
}

func NewTaskRamStorage(db *sql.DB) Task {
	return Task{
		db: db,
	}
}

func (s *Task) Get(id int) (*domain.Task, error) {
	rows, err := s.db.Query(`select * from todos where id=$1`, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {

	}

	return nil, repo.ErrNotImplemented
}

func (s *Task) GetPage(page int) ([]domain.Task, error) {
	return nil, repo.ErrNotImplemented
}

func (s *Task) Create(title, description string) (*domain.Task, error) {
	return nil, repo.ErrNotImplemented
}

func (s *Task) Delete(id int) error {
	return repo.ErrNotImplemented
}
