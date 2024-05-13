package rotator

import (
	"context"
	"fmt"
	"github.com/NoisyPunk/multiarmedbandit/internal/algorithm"
	rotatorconfig "github.com/NoisyPunk/multiarmedbandit/internal/configs"
	"github.com/NoisyPunk/multiarmedbandit/internal/storage"
	"github.com/google/uuid"
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

type Application interface {
	AddBanner(ctx context.Context, description string) (id uuid.UUID, err error)
	AddGroup(ctx context.Context, description string) (id uuid.UUID, err error)
	AddSlot(ctx context.Context, description string) (id uuid.UUID, err error)
	AddRotation(ctx context.Context, bannerId, slotId, groupId uuid.UUID) (id uuid.UUID, err error)
	ChooseBannerForSlot(ctx context.Context, slotId, groupId uuid.UUID) (bannerID uuid.UUID, err error)
	RegisterClick(ctx context.Context, rotationID uuid.UUID) (err error)
}

func (a App) AddBanner(ctx context.Context, description string) (id uuid.UUID, err error) {
	return a.Storage.AddBanner(ctx, description)
}

func (a App) AddGroup(ctx context.Context, description string) (id uuid.UUID, err error) {
	return a.Storage.AddGroup(ctx, description)
}

func (a App) AddSlot(ctx context.Context, description string) (id uuid.UUID, err error) {
	return a.Storage.AddSlot(ctx, description)
}

func (a App) AddRotation(ctx context.Context, bannerId, slotId, groupId uuid.UUID) (id uuid.UUID, err error) {
	return a.Storage.AddRotation(ctx, bannerId, slotId, groupId)
}

func (a App) ChooseBannerForSlot(ctx context.Context, slotId, groupId uuid.UUID) (bannerID uuid.UUID, err error) {
	rotations, err := a.Storage.GetRotationsForSlot(ctx, slotId, groupId)
	bestRotation, err := algorithm.ChooseBanner(rotations)
	if err != nil {
		return uuid.Nil, err
	}
	err = a.Storage.RegisterShown(ctx, bestRotation.Id)
	return bestRotation.BannerId, err
}

func (a App) RegisterClick(ctx context.Context, rotationID uuid.UUID) (err error) {
	return a.Storage.RegisterClick(ctx, rotationID)
}
