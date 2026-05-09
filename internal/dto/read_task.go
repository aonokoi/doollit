package dto

import (
	"proj/doollit/internal/domain"
)

type ReadTaskInput struct {
	ID int `json:"id"`
}

type ReadTaskOutput struct {
	domain.Task
}
