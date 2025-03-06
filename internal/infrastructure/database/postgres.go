package database

import (
	"context"
	"echo-template/internal/infrastructure"
	"echo-template/internal/infrastructure/logger"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresDB(config *infrastructure.Config, log *logger.Logger) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Name,
		config.Database.Password,
	)
	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf(
			"failed to connect to database with dsn: %s; exited with error: %s", dsn, err.Error(),
		)
		return nil, err
	}
	return db, nil
}
