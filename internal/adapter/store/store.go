package store

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/go-game-dev/rest-tm/internal/config"
)

type Store struct {
	DB *sqlx.DB
}    

func New(cfg config.Database) (*Store, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.Database,
	)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "sqlx.Open")
	}

	return &Store{
		DB: db,
	}, nil
}