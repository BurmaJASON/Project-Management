package project

type Service interface {
	CreateProject(project *Project) error
	GetProject(id int) (*Project, error)
	UpdateProject(project *Project) error
	DeleteProject(id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateProject(project *Project) error {
	return s.repo.Create(project)
}

func (s *service) GetProject(id int) (*Project, error) {
	return s.repo.GetByID(id)
}

func (s *service) UpdateProject(project *Project) error {
	return s.repo.Update(project)
}

func (s *service) DeleteProject(id int) error {
	return s.repo.Delete(id)
}
