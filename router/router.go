package router

import (
	"encoding/json"
	"managementapi/config"
	"managementapi/model"
	"net/http"

	"github.com/gorilla/mux"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var task model.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskConnect := config.Connection()

	result, err := taskConnect.CreateTask(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var task model.Task
	err := json.NewDecoder(r.Body).Decode(&task.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	taskConnect := config.Connection()
	result, err := taskConnect.UpdateTaskToOngoing(id, &task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}

func Tasks(w http.ResponseWriter, r *http.Request){
	taskConnect := config.Connection()
	result, err := taskConnect.AllTask()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
}

func Delete(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id := params["id"]
	taskConnect := config.Connection()
	result, err := taskConnect.DeleteTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
}