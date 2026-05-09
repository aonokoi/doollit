package usecase

import (
	"context"
	"fmt"

	"proj/doollit/internal/domain"
	"proj/doollit/internal/dto"
)

func (s *STask) CreateTask(ctx context.Context, input dto.CreateTaskInput) (dto.CreatTaskOutput, error) {
	const op = "usecase.CreateTask"

	var output dto.CreatTaskOutput

	task, err := domain.NewTask(input.Desc, input.CreatorName)
	if err != nil {
		return output, fmt.Errorf("new task: %s: %w", op, err)
	}

	id, err := s.postgres.CreateTask(ctx, task)
	if err != nil {
		return output, fmt.Errorf("saving to db: %s: %w", op, err)
	}

	return dto.CreatTaskOutput{
		ID: id,
	}, nil
}
