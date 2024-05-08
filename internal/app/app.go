package rotator

import (
	"context"
	"fmt"
	"github.com/NoisyPunk/multiarmedbandit/internal/configs"
	"github.com/NoisyPunk/multiarmedbandit/internal/storage"
)

type App struct {
	Storage storage.Storage
}

func New(ctx context.Context, config *rotatorconfig.Config) (*App, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DSN.Host, config.DSN.Port, config.DSN.User, config.DSN.Password, config.DSN.DBName, config.DSN.Ssl)

	store := storage.New()
	err := store.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &App{
		Storage: *store,
	}, nil
}
