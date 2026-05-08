package dto

import (
	"time"

	"proj/doollit/internal/domain"
)

type ReadTaskInput struct {
	ID int `json:"id"`
}

type ReadTaskOutput struct {
	ID          int                `json:"id"`
	Desc        domain.Description `json:"desc"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	CreatorName domain.Name        `json:"creator_name"`
}
