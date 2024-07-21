package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository for testing
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(task *Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockRepository) GetByID(id int) (*Task, error) {
	args := m.Called(id)
	return args.Get(0).(*Task), args.Error(1)
}

func (m *MockRepository) Update(task *Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Unit tests
func TestCreateTask(t *testing.T) {
	repo := new(MockRepository)
	service := NewService(repo)
	task := &Task{Name: "Test Task"}

	repo.On("Create", task).Return(nil)

	err := service.CreateTask(task)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestGetTask(t *testing.T) {
	repo := new(MockRepository)
	service := NewService(repo)
	task := &Task{ID: 1, Name: "Test Task"}

	repo.On("GetByID", 1).Return(task, nil)

	result, err := service.GetTask(1)
	assert.NoError(t, err)
	assert.Equal(t, task, result)
	repo.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
	repo := new(MockRepository)
	service := NewService(repo)
	task := &Task{ID: 1, Name: "Updated Task"}

	repo.On("Update", task).Return(nil)

	err := service.UpdateTask(task)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	repo := new(MockRepository)
	service := NewService(repo)

	repo.On("Delete", 1).Return(nil)

	err := service.DeleteTask(1)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}
