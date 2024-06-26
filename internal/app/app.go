package rotator

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/NoisyPunk/multiarmedbandit/internal/algorithm"
	rotatorconfig "github.com/NoisyPunk/multiarmedbandit/internal/configs"
	"github.com/NoisyPunk/multiarmedbandit/internal/logger"
	"github.com/NoisyPunk/multiarmedbandit/internal/queue"
	"github.com/NoisyPunk/multiarmedbandit/internal/storage"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type App struct {
	Storage  storage.Storage
	Producer *queue.Producer
}

func New(ctx context.Context, config *rotatorconfig.Config) (*App, error) {
	l := logger.FromContext(ctx)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DSN.Host, config.DSN.Port, config.DSN.User, config.DSN.Password, config.DSN.DBName, config.DSN.Ssl)

	store := storage.New()
	err := store.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	producer, err := queue.NewProducer(l, config)
	if err != nil {
		return nil, err
	}

	return &App{
		Storage:  *store,
		Producer: producer,
	}, nil
}

type Application interface {
	AddBanner(ctx context.Context, description string) (id uuid.UUID, err error)
	AddGroup(ctx context.Context, description string) (id uuid.UUID, err error)
	AddSlot(ctx context.Context, description string) (id uuid.UUID, err error)
	AddRotation(ctx context.Context, bannerID, slotID, groupID string) (id uuid.UUID, err error)
	ChooseRotationForSlot(ctx context.Context, slotID, groupID string) (rotation storage.Rotation, err error)
	RegisterClick(ctx context.Context, rotationID string) (err error)
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

func (a App) AddRotation(ctx context.Context, bannerID, slotID, groupID string) (id uuid.UUID, err error) {
	bannerUUID, err := uuid.Parse(bannerID)
	if err != nil {
		return uuid.Nil, err
	}
	slotUUID, err := uuid.Parse(slotID)
	if err != nil {
		return uuid.Nil, err
	}
	groupUUID, err := uuid.Parse(groupID)
	if err != nil {
		return uuid.Nil, err
	}
	return a.Storage.AddRotation(ctx, bannerUUID, slotUUID, groupUUID)
}

func (a App) ChooseRotationForSlot(ctx context.Context, slotID, groupID string) (rotation storage.Rotation, err error) {
	slotUUID, err := uuid.Parse(slotID)
	if err != nil {
		return rotation, err
	}
	groupUUID, err := uuid.Parse(groupID)
	if err != nil {
		return rotation, err
	}
	rotations, err := a.Storage.GetSlotRotations(ctx, slotUUID, groupUUID)
	if err != nil {
		return rotation, err
	}
	bestRotation, err := algorithm.ChooseBanner(rotations)
	if err != nil {
		return rotation, err
	}
	err = a.Storage.RegisterShown(ctx, bestRotation.ID)
	if err != nil {
		return rotation, err
	}
	err = a.publishShown(ctx, bestRotation)
	return bestRotation, err
}

func (a App) RegisterClick(ctx context.Context, rotationID string) (err error) {
	rotationUUID, err := uuid.Parse(rotationID)
	if err != nil {
		return err
	}
	err = a.publishClick(ctx, rotationUUID)
	if err != nil {
		return err
	}
	return a.Storage.RegisterClick(ctx, rotationUUID)
}

func (a App) publishClick(ctx context.Context, rotationID uuid.UUID) (err error) {
	l := logger.FromContext(ctx)

	rotation, err := a.Storage.GetRotation(ctx, rotationID)
	if err != nil {
		return err
	}

	message := queue.Event{
		Name:        "Show",
		SlotID:      rotation.SlotID.String(),
		BannerID:    rotation.BannerID.String(),
		GroupID:     rotation.GroupID.String(),
		DateAndTime: time.Now(),
	}

	j, err := json.Marshal(message)
	if err != nil {
		l.Error("can't marshal event for queue", zap.String("error_message", err.Error()))
	}

	err = a.Producer.RmqChannel.Publish(
		"",
		"CalendarQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "json",
			Body:        j,
		},
	)
	return err
}

func (a App) publishShown(ctx context.Context, rotation storage.Rotation) (err error) {
	l := logger.FromContext(ctx)

	message := queue.Event{
		Name:        "Show",
		SlotID:      rotation.SlotID.String(),
		BannerID:    rotation.BannerID.String(),
		GroupID:     rotation.GroupID.String(),
		DateAndTime: time.Now(),
	}

	j, err := json.Marshal(message)
	if err != nil {
		l.Error("can't marshal event for queue", zap.String("error_message", err.Error()))
	}

	err = a.Producer.RmqChannel.Publish(
		"",
		"CalendarQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "json",
			Body:        j,
		},
	)
	return err
}
