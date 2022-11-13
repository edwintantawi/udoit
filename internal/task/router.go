package task

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(path string, mux *mux.Router, handler Handler) {
	r := mux.PathPrefix(path).Subrouter()
	r.Methods(http.MethodGet).Path("/").HandlerFunc(handler.GetAllTask)
	r.Methods(http.MethodPost).Path("/").HandlerFunc(handler.PostNewTask)
	r.Methods(http.MethodGet).Path("/{taskId}").HandlerFunc(handler.GetTaskByID)
	r.Methods(http.MethodPut).Path("/{taskId}").HandlerFunc(handler.PutTaskByID)
	r.Methods(http.MethodDelete).Path("/{taskId}").HandlerFunc(handler.DeleteTaskByID)
}
