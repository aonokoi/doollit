package postgres

import (
	"context"
	"fmt"

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
		"postgres://%s:%s@%s:%s/%s",
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
	id := 1
	return id, nil
}

func (p *Pool) GetTask(ctx context.Context, id int) (domain.Task, error) {
	var task domain.Task

	return task, nil
}

func (p *Pool) DeleteTask(ctx context.Context, id int) error {
	return nil
}

func (p *Pool) Close() {
	// shutdown
}
