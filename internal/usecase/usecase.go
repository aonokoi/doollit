package usecase

import (
	"context"

	"proj/doollit/internal/adapter/postgres"
	"proj/doollit/internal/domain"
)

type Postgres interface {
	CreateTask(ctx context.Context, task domain.Task) (int, error)
	GetTask(ctx context.Context, id int) (domain.Task, error)
	DeleteTask(ctx context.Context, id int) error
	// UpdateTask(ctx context.Context, id int,)
}

type STask struct {
	postgres Postgres
}

func NewSTask(postgres *postgres.Pool) *STask {
	return &STask{
		postgres: postgres,
	}
}
