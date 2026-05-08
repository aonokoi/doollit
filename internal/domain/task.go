package domain

import "time"

type Description string

type Name string

type Task struct {
	ID          int         `json:"id"`
	Desc        Description `json:"desc"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	CreatorName Name        `json:"creator_name"`
}

func NewTask(desc, CreatorName string) (Task, error) {
	now := time.Now()
	task := Task{
		Desc:        Description(desc),
		CreatorName: Name(CreatorName),
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// validate

	return task, nil
}
