package task

import (
	"encoding/json"
	"net/http"

	"github.com/edwintantawi/udoit/internal/errorx"
	"github.com/edwintantawi/udoit/internal/response"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler interface {
	PostNewTask(w http.ResponseWriter, r *http.Request)
	GetAllTask(w http.ResponseWriter, r *http.Request)
	GetTaskByID(w http.ResponseWriter, r *http.Request)
	PutTaskByID(w http.ResponseWriter, r *http.Request)
	DeleteTaskByID(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	taskService Service
}

func NewHandler(taskService Service) *handler {
	return &handler{taskService: taskService}
}

func (h *handler) PostNewTask(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w)

	var taskIn TaskIn
	if err := json.NewDecoder(r.Body).Decode(&taskIn); err != nil {
		resp.Fail(errorx.NewInvariant(err.Error()), nil)
		return
	}

	validate := validator.New()
	if err := validate.Struct(taskIn); err != nil {
		resp.Fail(err, nil)
		return
	}

	err := h.taskService.Add(r.Context(), taskIn)
	if err != nil {
		resp.Fail(err, nil)
		return
	}

	resp.Success(http.StatusCreated, "Successfully added new task", nil, nil)
}

func (h *handler) GetAllTask(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w)

	result, err := h.taskService.GetAll(r.Context())
	if err != nil {
		resp.Fail(err, nil)
		return
	}

	code := http.StatusOK
	resp.Success(http.StatusOK, http.StatusText(code), nil, result)
}

func (h *handler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w)

	id := ID(mux.Vars(r)["taskId"])

	result, err := h.taskService.GetByID(r.Context(), id)
	if err != nil {
		resp.Fail(err, nil)
		return
	}

	code := http.StatusOK
	resp.Success(http.StatusOK, http.StatusText(code), nil, result)
}

func (h *handler) PutTaskByID(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w)

	id := ID(mux.Vars(r)["taskId"])
	var taskIn TaskIn
	if err := json.NewDecoder(r.Body).Decode(&taskIn); err != nil {
		resp.Fail(errorx.NewInvariant(err.Error()), nil)
		return
	}

	err := h.taskService.UpdateByID(r.Context(), id, taskIn)
	if err != nil {
		resp.Fail(err, nil)
		return
	}

	resp.Success(http.StatusOK, "Successfully changed task", nil, nil)
}

func (h *handler) DeleteTaskByID(w http.ResponseWriter, r *http.Request) {
	resp := response.New(w)

	id := ID(mux.Vars(r)["taskId"])

	err := h.taskService.DeleteByID(r.Context(), id)
	if err != nil {
		resp.Fail(err, nil)
		return
	}

	resp.Success(http.StatusOK, "Successfully deleted task", nil, nil)
}
