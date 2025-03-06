package application

import (
	"echo-template/internal/delivery/rest/router"
	"echo-template/internal/infrastructure"
	"echo-template/internal/infrastructure/database"
	"echo-template/internal/infrastructure/logger"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	tokenjwt "echo-template/pkg/token_jwt"
)

type Application struct {
	e       *echo.Echo
	Address string
	Db      *pgxpool.Pool
	Logger  *logger.Logger
}

func NewApplication(config *infrastructure.Config) *Application {
	e := echo.New()
	l := logger.NewLogger() // Создаём логгер

	db, err := database.NewPostgresDB(config, l)
	if err != nil {
		l.Errorf("failed to connect to database: %s", err.Error())
		return nil
	}

	tokenjwt.InitJWTKey(config.Other.JWTKey)

	return &Application{
		e:       e,
		Address: config.Server.Address,
		Db:      db,
		Logger:  l, // Добавляем логгер
	}
}

func (a *Application) RunServer() error {
	e := initServer(a)
	a.Logger.Info("Starting server on " + a.Address)

	if err := e.Start(a.Address); err != nil {
		if err != http.ErrServerClosed {
			a.Logger.Errorf("Failed to start server: %s", err.Error())
		}
	}
	return nil
}

func initServer(a *Application) *echo.Echo {
	e := a.e
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Теперь логгер передаётся корректно
	router.RegisterRouter(e, a.Db, a.Logger)

	return e
}
