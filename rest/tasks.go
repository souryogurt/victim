package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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

func GetAllTasks(svc TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tasks, err := svc.GetAllTasks(ctx)
		if err != nil {
			svc.Println(ctx, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		js, err := json.Marshal(tasks)
		if err != nil {
			svc.Println(ctx, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(js); err != nil {
			svc.Println(ctx, err)
		}
	}
}

func CreateTask(svc TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var task *victim.Task
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&task); err != nil {
			svc.Println(ctx, err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		resp, err := svc.CreateTask(ctx, task)
		if err != nil {
			svc.Println(ctx, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		js, err := json.Marshal(resp)
		if err != nil {
			svc.Println(ctx, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(js); err != nil {
			svc.Println(ctx, err)
		}
	}
}

func GetTask(svc TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ID, err := strconv.Atoi(chi.URLParam(r, "taskID"))
		if err != nil {
			svc.Println(ctx, err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		task, err := svc.GetTask(ctx, ID)
		if err != nil {
			svc.Println(ctx, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		js, err := json.Marshal(task)
		if err != nil {
			svc.Println(ctx, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(js); err != nil {
			svc.Println(ctx, err)
		}
	}
}

func UpdateTask(svc TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var task *victim.Task
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&task); err != nil {
			svc.Println(ctx, err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		resp, err := svc.UpdateTask(ctx, task)
		if err != nil {
			svc.Println(ctx, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		js, err := json.Marshal(resp)
		if err != nil {
			svc.Println(ctx, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(js); err != nil {
			svc.Println(ctx, err)
		}
	}
}

func DeleteTask(svc TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ID, err := strconv.Atoi(chi.URLParam(r, "taskID"))
		if err != nil {
			svc.Println(ctx, err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		taskID, err := svc.DeleteTask(ctx, ID)
		if err != nil {
			svc.Println(ctx, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		js, err := json.Marshal(taskID)
		if err != nil {
			svc.Println(ctx, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(js); err != nil {
			svc.Println(ctx, err)
		}
	}
}
