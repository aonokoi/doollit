package dto

import "proj/doollit/internal/domain"

type TaskUpdate struct {
	Desc *domain.Description `json:"desc,omitempty"`
}

type UpdateTaskInput struct {
	ID int `json:"id"`
	TaskUpdate
}

type UpdateTaskOutput struct{}
