package usecase

import (
	"context"
	"fmt"

	"proj/doollit/internal/domain"
	"proj/doollit/internal/dto"
)

func (s *STask) CreateTask(ctx context.Context, input dto.CreateTaskInput) (dto.CreatTaskOutput, error) {
	var output dto.CreatTaskOutput

	task, err := domain.NewTask(input.Desc, input.CreatedName)
	if err != nil {
		return output, fmt.Errorf("new task: %w", err)
	}

	id, err := s.postgres.CreateTask(ctx, task)
	if err != nil {
		return output, fmt.Errorf("saving to db: %w", err)
	}

	output.ID = id
	return output, nil
}
