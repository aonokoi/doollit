package http

import "proj/doollit/internal/usecase"

type Handlers struct {
	taskService *usecase.STask
}

func NewHandlers(taskService *usecase.STask) *Handlers {
	return &Handlers{
		taskService: taskService,
	}
}
