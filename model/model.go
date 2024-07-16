package model

import (
	

	"go.mongodb.org/mongo-driver/bson/primitive"
	
)

type Status struct {
	Done       bool `json:"done"`
	Ongoing    bool `json:"ongoing"`
	NotStarted bool `json:"not_started"`
}

type Task struct {
	ID       primitive.ObjectID `json:"task_id" bson:"_id,omitempty"`
	Name     string             `json:"name"`
	Content  string             `json:"content"`
	Deadline string             `json:"deadline"`
	Status   Status             `json:"track"`
}

func (t *Task) Create(task Task) {


}
