package ram

import (
	"sort"

	"todo-golang.com/domain"
	"todo-golang.com/repo"
)

type Task struct {
	data    map[int]domain.Task
	counter int
}

func NewTaskRamStorage() Task {
	return Task{
		data: map[int]domain.Task{
			0: domain.NewTask(0, "1", "", false),
			1: domain.NewTask(1, "2", "", false),
			2: domain.NewTask(2, "3", "", false),
			3: domain.NewTask(3, "4", "", false),
		},
		counter: 4,
	}
}

func (s *Task) Get(id int) (*domain.Task, error) {
	if task, ok := s.data[id]; ok {
		return &task, nil
	}

	return nil, repo.ErrNotFound
}

func (s *Task) GetPage(page int) ([]domain.Task, error) {
	count := min(len(s.data)-page*repo.PAGE_SIZE, repo.PAGE_SIZE)

	if count <= 0 {
		return nil, repo.ErrOutOfRange
	}

	c1, c2 := page, 0
	result := make([]domain.Task, count)

	for id := range s.data {
		if c1 > 0 {
			c1 -= 1
			continue
		}
		if c2 < count {
			result[c2] = s.data[id]
			c2++
			continue
		}
		break
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Id < result[j].Id
	})

	return result, nil
}

func (s *Task) Create(title, description string) (*domain.Task, error) {
	newTask := domain.Task{
		Id:          s.counter,
		Title:       title,
		Description: description,
		IsDone:      false,
	}

	s.counter += 1

	s.data[newTask.Id] = newTask

	return &newTask, nil
}

func (s *Task) Delete(id int) error {
	if _, ok := s.data[id]; !ok {
		return repo.ErrNotFound
	}

	delete(s.data, id)
	return nil
}
