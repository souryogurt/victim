package victim

import (
	"context"
	"log"

	"github.com/pkg/errors"
)

// Task represents single task in todo list
// swagger:model
type Task struct {
	// The ID of the task
	// read only: true
	ID int `json:"id"`

	// The description of the task
	// required: true
	// example: do not forget to buy a milk
	Text string `json:"text"`

	// The status of task completion
	// required: true
	// example: false
	IsCompleted bool `json:"completed"`
}

type Victim struct {
}

func NewVictim() *Victim {
	return &Victim{}
}

func (svc *Victim) GetAllTasks(ctx context.Context) ([]*Task, error) {
	return nil, errors.New("Not Implemented")
}

func (svc *Victim) CreateTask(ctx context.Context, task *Task) (*Task, error) {
	return nil, errors.New("Not Implemented")
}

func (svc *Victim) GetTask(ctx context.Context, id int) (*Task, error) {
	return nil, errors.New("Not Implemented")
}

func (svc *Victim) UpdateTask(ctx context.Context, task *Task) (*Task, error) {
	return nil, errors.New("Not Implemented")
}

func (svc *Victim) DeleteTask(ctx context.Context, id int) (int, error) {
	return 0, errors.New("Not Implemented")
}

func (svc *Victim) Println(ctx context.Context, v ...interface{}) {
	log.Println(v...)
}
