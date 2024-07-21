package task

import "database/sql"

type Repository interface {
	Create(task *Task) error
	GetByID(id int) (*Task, error)
	Update(project *Task) error
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(task *Task) error {
	_, err := r.db.Exec("INSERT INTO tasks (project_id, name) VALURES (?, ?)", task.ProjectId, task.Name)

	return err
}

func (r *repository) GetByID(id int) (*Task, error) {
	row := r.db.QueryRow("SELECT id, project_id, name from tasks WHERE id = ?", id)

	task := &Task{}
	err := row.Scan(&task.ID, &task.Name, &task.ProjectId)

	return task, err
}

func (r *repository) Update(task *Task) error {
	_, err := r.db.Exec("UPDATE tasks SET project_id = ? name = ? WHERE id =? ", task.ProjectId, task.Name, task.ID)

	return err
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FORM tasks WHERE id = ?", id)

	return err
}
