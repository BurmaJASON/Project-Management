package task

type Task struct {
	ID        int    `json:"id"`
	ProjectId int    `json:"project_id"`
	Name      string `json:"name"`
}
