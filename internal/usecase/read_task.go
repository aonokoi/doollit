package usecase

import (
	"context"
	"fmt"

	"proj/doollit/internal/dto"
)

func (s *STask) ReadTask(ctx context.Context, input dto.ReadTaskInput) (dto.ReadTaskOutput, error) {
	const op = "usecase.Read"

	var output dto.ReadTaskOutput

	task, err := s.postgres.ReadTask(ctx, input.ID)
	if err != nil {
		return output, fmt.Errorf("unable to read task: %s: %w", op, err)
	}

	return dto.ReadTaskOutput{
		ID:          task.ID,
		Desc:        task.Desc,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
		CreatorName: task.CreatorName,
	}, nil
}
