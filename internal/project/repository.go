package project

import (
	"database/sql"
)

type Repository interface {
	Create(project *Project) error
	GetByID(id int) (*Project, error)
	Update(project *Project) error
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(project *Project) error {
	_, err := r.db.Exec("INSERT INTO projects (name) VALUES (?)", project.Name)
	return err
}

func (r *repository) GetByID(id int) (*Project, error) {
	row := r.db.QueryRow("SELECT id, name from projects WHERE id =? ", id)

	project := &Project{}
	err := row.Scan(&project.ID, &project.Name)

	return project, err
}

func (r *repository) Update(project *Project) error {
	_, err := r.db.Exec("UPDATE projects SET name = ? WHERE id = ?", project.Name, project.ID)
	return err
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM projects WHERE id = ?", id)

	return err
}
