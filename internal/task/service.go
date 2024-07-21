package task

type Service interface {
	CreateTask(task *Task) error
	GetTask(id int) (*Task, error)
	UpdateTask(task *Task) error
	DeleteTask(id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateTask(task *Task) error {
	return s.repo.Create(task)
}

func (s *service) GetTask(id int) (*Task, error) {
	return s.repo.GetByID(id)
}

func (s *service) UpdateTask(task *Task) error {
	return s.repo.Update(task)
}

func (s *service) DeleteTask(id int) error {
	return s.repo.Delete(id)
}
