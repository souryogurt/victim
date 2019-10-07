// REST API for TODO application
//
// Provides REST API for create, read, update and delete tasks.
//
//     Schemes: http
//     BasePath: /
//     Version: 0.0.1
//     Host: localhost
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/souryogurt/victim"
)

type TaskService interface {
	GetAllTasks(ctx context.Context) ([]*victim.Task, error)
	CreateTask(ctx context.Context, task *victim.Task) (*victim.Task, error)
	GetTask(ctx context.Context, ID int) (*victim.Task, error)
	UpdateTask(ctx context.Context, task *victim.Task) (*victim.Task, error)
	DeleteTask(ctx context.Context, ID int) (int, error)
	Println(ctx context.Context, v ...interface{})
}

// A list of all requested tasks
// swagger:response getAllTasksResponse
type GetAllTasksResponse struct {
	// The array of all tasks
	// in: body
	Body []*victim.Task
}

// swagger:route GET /tasks tasks getalltasks
// List all tasks
//
// List all tasks
//     Responses:
//	 200: getAllTasksResponse
//       500: description:Internal server error
func GetAllTasks(svc TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tasks, err := svc.GetAllTasks(ctx)
		if err != nil {
			err = errors.Wrap(err, "can't retreive tasks")
			svc.Println(ctx, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		js, err := json.Marshal(tasks)
		if err != nil {
			err = errors.Wrap(err, "can't marshal response")
			svc.Println(ctx, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(js); err != nil {
			err = errors.Wrap(err, "can't write response")
			svc.Println(ctx, err)
		}
	}
}

// swagger:parameters createtask
type CreateTaskParams struct {
	// Required: true
	// in: body
	Body victim.Task
}

// An attributes of created task
// swagger:response createTaskResponse
type CreateTaskResponse struct {
	// Created task parameters
	//
	// in: body
	Body *victim.Task
}

// swagger:route POST /tasks tasks createtask
// Create new task
//
// Create new task
//     Responses:
//	 201: createTaskResponse
//	 400: description:Bad incoming JSON
//       500: description:Internal server error
func CreateTask(svc TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var task *victim.Task
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&task); err != nil {
			err = errors.Wrap(err, "can't decode request payload")
			svc.Println(ctx, err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		resp, err := svc.CreateTask(ctx, task)
		if err != nil {
			err = errors.Wrap(err, "can't create task")
			svc.Println(ctx, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		js, err := json.Marshal(resp)
		if err != nil {
			err = errors.Wrap(err, "can't marshal response")
			svc.Println(ctx, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if _, err := w.Write(js); err != nil {
			err = errors.Wrap(err, "can't write response")
			svc.Println(ctx, err)
		}
	}
}

// swagger:parameters gettask
type GetTaskParams struct {
	// Identifier of the task
	// Required: true
	// in: path
	ID int
}

// A requested task
// swagger:response getTaskResponse
type GetTaskResponse struct {
	// The task
	// in: body
	Body *victim.Task
}

// swagger:route GET /task/{ID} tasks gettask
// Get task by ID
//
// Get task by ID
//     Responses:
//	 200: getTaskResponse
//       400: description:Invalid task ID
//       500: description:Internal server error
func GetTask(svc TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ID, err := strconv.Atoi(chi.URLParam(r, "taskID"))
		if err != nil {
			err = errors.Wrap(err, "can't parse task ID")
			svc.Println(ctx, err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		task, err := svc.GetTask(ctx, ID)
		if err != nil {
			err = errors.Wrap(err, "can't retreive task")
			svc.Println(ctx, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		js, err := json.Marshal(task)
		if err != nil {
			err = errors.Wrap(err, "can't marshal response")
			svc.Println(ctx, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(js); err != nil {
			err = errors.Wrap(err, "can't write response")
			svc.Println(ctx, err)
		}
	}
}

// swagger:parameters updatetask
type UpdateTaskParams struct {
	// Identifier of the task
	// Required: true
	// in: path
	ID int
}

// An updated task attributes
// swagger:response updateTaskResponse
type UpdateTaskResponse struct {
	// Updated task attributes
	// in: body
	Body *victim.Task
}

// swagger:route PUT /task/{ID} tasks updatetask
// Update task
//
// Update task
//     Responses:
//	 200: updateTaskResponse
//       400: description:Invalid task ID
//       500: description:Internal server error
func UpdateTask(svc TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ID, err := strconv.Atoi(chi.URLParam(r, "taskID"))
		if err != nil {
			err = errors.Wrap(err, "can't parse task ID")
			svc.Println(ctx, err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		var task *victim.Task
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&task); err != nil {
			err = errors.Wrap(err, "can't decode request payload")
			svc.Println(ctx, err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		task.ID = ID
		resp, err := svc.UpdateTask(ctx, task)
		if err != nil {
			err = errors.Wrap(err, "can't update task")
			svc.Println(ctx, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		js, err := json.Marshal(resp)
		if err != nil {
			err = errors.Wrap(err, "can't marshal response")
			svc.Println(ctx, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(js); err != nil {
			err = errors.Wrap(err, "can't write response")
			svc.Println(ctx, err)
		}
	}
}

// swagger:parameters deletetask
type DeleteTaskParams struct {
	// Identifier of the task
	// Required: true
	// in: path
	ID int
}

// An ID of deleted task
// swagger:response deleteTaskResponse
type DeleteTaskResponse struct {
	// Deleted task ID
	// in: body
	Body int
}

// swagger:route DELETE /task/{ID} tasks deletetask
// Delete task
//
// Delete task
//     Responses:
//	 200: deleteTaskResponse
//       400: description:Invalid task ID
//       500: description:Internal server error
func DeleteTask(svc TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ID, err := strconv.Atoi(chi.URLParam(r, "taskID"))
		if err != nil {
			err = errors.Wrap(err, "can't parse task ID")
			svc.Println(ctx, err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		taskID, err := svc.DeleteTask(ctx, ID)
		if err != nil {
			err = errors.Wrap(err, "can't delete task")
			svc.Println(ctx, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		js, err := json.Marshal(taskID)
		if err != nil {
			err = errors.Wrap(err, "can't marshal response")
			svc.Println(ctx, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(js); err != nil {
			err = errors.Wrap(err, "can't write response")
			svc.Println(ctx, err)
		}
	}
}
