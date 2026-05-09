package usecase

import (
	"context"
	"fmt"

	"proj/doollit/internal/dto"
)

func (s *STask) UpdateTask(ctx context.Context, input dto.UpdateTaskInput) (dto.UpdateTaskOutput, error) {
	const op = "usecase.UpdateTask"

	var output dto.UpdateTaskOutput

	err := s.postgres.UpdateTask(ctx, input.ID, *input.Desc)
	if err != nil {
		return output, fmt.Errorf("unable to update task: %s: %w", op, err)
	}

	return output, nil
}
