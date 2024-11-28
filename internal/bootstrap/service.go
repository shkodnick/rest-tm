package bootstrap

import (
	"github.com/go-game-dev/rest-tm/internal/config"
	tm_store "github.com/go-game-dev/rest-tm/internal/adapter/store/tm"
	"github.com/go-game-dev/rest-tm/internal/adapter/store"
	"github.com/pkg/errors"
	"github.com/go-game-dev/rest-tm/internal/service/tm"
)

func CreateTmService() (*tm.TmService, error) {
	cfg, err := config.NewConfig()
	if err != nil {
			return nil, errors.Wrap(err, "CreateConfig")
	}

	dbStore, err := store.New(cfg.Database)
	if err != nil {
			return nil, err
	}

	tmStore := &tm_store.TmStore{Store: dbStore}
	return tm.New(tmStore)
}