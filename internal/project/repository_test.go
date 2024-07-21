package project

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository for testing
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(project *Project) error {
	args := m.Called(project)
	return args.Error(0)
}

func (m *MockRepository) GetByID(id int) (*Project, error) {
	args := m.Called(id)
	return args.Get(0).(*Project), args.Error(1)
}

func (m *MockRepository) Update(project *Project) error {
	args := m.Called(project)
	return args.Error(0)
}

func (m *MockRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Unit tests
func TestCreateProject(t *testing.T) {
	repo := new(MockRepository)
	service := NewService(repo)
	project := &Project{Name: "Test Project"}

	repo.On("Create", project).Return(nil)

	err := service.CreateProject(project)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestGetProject(t *testing.T) {
	repo := new(MockRepository)
	service := NewService(repo)
	project := &Project{ID: 1, Name: "Test Project"}

	repo.On("GetByID", 1).Return(project, nil)

	result, err := service.GetProject(1)
	assert.NoError(t, err)
	assert.Equal(t, project, result)
	repo.AssertExpectations(t)
}

func TestUpdateProject(t *testing.T) {
	repo := new(MockRepository)
	service := NewService(repo)
	project := &Project{ID: 1, Name: "Updated Project"}

	repo.On("Update", project).Return(nil)

	err := service.UpdateProject(project)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestDeleteProject(t *testing.T) {
	repo := new(MockRepository)
	service := NewService(repo)

	repo.On("Delete", 1).Return(nil)

	err := service.DeleteProject(1)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}
