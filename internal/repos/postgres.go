package repos

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

const listsTable = "tasks"

func NewPostgresDB(ctx context.Context, cfg Config) (*pgx.Conn, error) {
	db, err := pgx.Connect(ctx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode))
	if err != nil {
		slog.Error("error connecting to postgres", err.Error())
		return nil, err
	}

	if err = db.Ping(ctx); err != nil {
		slog.Error("error to ping postgres", err.Error())
		return nil, err
	}

	return db, nil
}
