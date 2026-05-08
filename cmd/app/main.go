package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"proj/doollit/config"
	"proj/doollit/internal/adapter/postgres"
	"proj/doollit/internal/usecase"
)

func main() {
	// init config
	c, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	// run app
	err = AppRun(context.Background(), c)
	if err != nil {
		panic(err)
	}
}

func AppRun(ctx context.Context, c *config.Config) error {
	// init postgres
	pgPool, err := postgres.New(ctx, c.Postgres)
	if err != nil {
		return fmt.Errorf("Unable to init postgres: %w", err)
	}

	// init usecase
	TaskUseCase := usecase.NewSTask(pgPool)
	_ = TaskUseCase
	// init router

	// init httpserver

	// Приложение запущено и готово к работе

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	<-sig // ждём здесь сигнала (Ctrl+C или SIGTERM)

	// Закрываем ресурсы
	pgPool.Close()

	return nil
}
