package usecase

import (
	"context"
	"fmt"

	"proj/doollit/internal/dto"
)

func (s *STask) DeleteTask(
	ctx context.Context, input dto.DeleteTaskInput,
) (dto.DeleteTaskOutput, error) {
	const op = "usecase.DeleteTask"

	var output dto.DeleteTaskOutput

	err := s.postgres.DeleteTask(ctx, input.ID)
	if err != nil {
		return output, fmt.Errorf("unable to delete task: %s: %w", op, err)
	}

	return output, nil
}
