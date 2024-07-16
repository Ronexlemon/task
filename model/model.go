package model

type Task struct {
	Id       int    `json:"task_id"`
	Name     string `json:"name"`
	Content  string `json:"content"`
	Deadline string `json:"deadline"`
	Status   Status `json:"isDone"`
}

type Status struct {
	Done       bool
	Ongoing    bool
	NotStarted bool
}

func (t *Task) Create(task Task) {
	

}
