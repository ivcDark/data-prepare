package app

import (
	"data-preparer/internal/config"
	"data-preparer/internal/handler"
	"data-preparer/internal/repository"
	"data-preparer/internal/service"
)

type App struct {
	cfg *config.Config
}

func New() (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	// Инициализация репозитория
	repo, err := repository.NewMySQL(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	// Инициализация сервисов
	leagueService := service.NewLeagueTableService(repo)
	playerService := service.NewPlayerStatsService(repo)

	// Инициализация обработчика
	handler, err := handler.NewRabbitMQHandler(
		cfg.RabbitMQURL,
		leagueService,
		playerService,
	)
	if err != nil {
		return nil, err
	}

	go handler.Start()

	return &App{cfg: cfg}, nil
}

func (a *App) Run() error {
	// Блокируем главный поток
	select {}
}