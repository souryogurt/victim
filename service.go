package victim

import (
	"context"
	"errors"
	"log"
)

type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
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
