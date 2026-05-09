package postgres

import (
	"context"
	"fmt"
	"time"

	"proj/doollit/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	User     string `envconfig:"DB_USER"`
	Password string `envconfig:"DB_PASSWORD"`
	Port     string `envconfig:"DB_PORT"`
	Host     string `envconfig:"DB_HOST"`
	DBName   string `envconfig:"DB_NAME"`
}

type Pool struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, c Config) (*Pool, error) {
	const op = "postgres.New"

	DBURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=5",
		c.User, c.Password, c.Host, c.Port, c.DBName)

	pool, err := pgxpool.New(ctx, DBURL)
	if err != nil {
		return nil, fmt.Errorf("ubable to connect to db: %s, %w", op, err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping DB: %s, %w", op, err)
	}

	return &Pool{pool: pool}, nil
}

func (p *Pool) CreateTask(ctx context.Context, task domain.Task) (int, error) {
	const op = "postgres.CreateTask"

	sql := `
	INSERT INTO tasks("desc", creator_name)
	VALUES($1, $2)
	RETURNING id
	`

	var id int

	err := p.pool.QueryRow(ctx, sql, task.Desc, task.CreatorName).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("unable to query db: %s: %w", op, err)
	}

	return id, nil
}

func (p *Pool) ReadTask(ctx context.Context, id int) (domain.Task, error) {
	const op = "postgres.ReadTask"

	sql := `SELECT * FROM tasks WHERE id = $1`

	var task domain.Task

	err := p.pool.QueryRow(ctx, sql, id).
		Scan(
			&task.ID,
			&task.Desc,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.CreatorName)
	if err != nil {
		return task, fmt.Errorf("unable to read the task: %s: %w", op, err)
	}

	return task, nil
}

func (p *Pool) DeleteTask(ctx context.Context, id int) error {
	const op = "postgres.DeleteTask"

	sql := `DELETE * FROM tasks WHERE id = $1`

	tag, err := p.pool.Exec(ctx, sql, id)
	if err != nil {
		return fmt.Errorf("unable to delete task: %s: %w", op, err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("task is not found: tag: %s: %d", op, tag.RowsAffected())
	}

	return nil
}

func (p *Pool) UpdateTask(ctx context.Context, id int, desc domain.Description) error {
	const op = "postgres.UpdateTask"
	now := time.Now()

	sql := `
	UPDATE tasks
	SET "desc" = $2,
		updated_at = $3
	WHERE id = $1
	`

	tag, err := p.pool.Exec(ctx, sql, id, desc, now)
	if err != nil {
		return fmt.Errorf("unable to update the task: %s: %w", op, err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("task in not found: tag: %s: %d", op, err)
	}

	return nil
}

func (p *Pool) Close() {
	// shutdown
}
