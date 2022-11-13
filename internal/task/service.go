package task

import (
	"context"
)

type TaskIn struct {
	Content     string `json:"content" validate:"required"`
	Description string `json:"description"`
}

type Service interface {
	GetAll(ctx context.Context) ([]Task, error)
	Add(ctx context.Context, v TaskIn) error
	GetByID(ctx context.Context, id ID) (Task, error)
	UpdateByID(ctx context.Context, id ID, v TaskIn) error
	DeleteByID(ctx context.Context, id ID) error
}

type service struct {
	taskRepo Repo
}

func NewService(taskRepo Repo) *service {
	return &service{taskRepo: taskRepo}
}

func (s *service) GetAll(ctx context.Context) ([]Task, error) {
	return s.taskRepo.FindAll(ctx)
}

func (s *service) Add(ctx context.Context, v TaskIn) error {
	return s.taskRepo.Create(ctx, v)
}

func (s *service) GetByID(ctx context.Context, id ID) (Task, error) {
	return s.taskRepo.FindByID(ctx, id)
}

func (s *service) UpdateByID(ctx context.Context, id ID, v TaskIn) error {
	return s.taskRepo.UpdateByID(ctx, id, v)
}

func (s *service) DeleteByID(ctx context.Context, id ID) error {
	return s.taskRepo.DeleteByID(ctx, id)
}
