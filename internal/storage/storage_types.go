package storage

import (
	"context"

	"github.com/google/uuid"
)

type Banner struct {
	ID          uuid.UUID `db:"id"`
	Description string    `db:"description"`
}

type Group struct {
	ID          uuid.UUID `db:"id"`
	Description string    `db:"description"`
}

type Slot struct {
	ID          uuid.UUID `db:"id"`
	Description string    `db:"description"`
}

type Rotation struct {
	ID       uuid.UUID `db:"id"`
	BannerID uuid.UUID `db:"banner_id"`
	GroupID  uuid.UUID `db:"group_id"`
	SlotID   uuid.UUID `db:"slot_id"`
	Clicks   int       `db:"clicks"`
	Shows    int       `db:"shows"`
}

type Storager interface {
	Connect(ctx context.Context, dsn string) (err error)
	Close() error
	AddBanner(ctx context.Context, description string) (id uuid.UUID, err error)

	AddGroup(ctx context.Context, description string) (id uuid.UUID, err error)

	AddSlot(ctx context.Context, description string) (id uuid.UUID, err error)

	AddRotation(ctx context.Context, bannerID, slotID, groupID uuid.UUID) (id uuid.UUID, err error)
	GetSlotRotations(ctx context.Context, slotID, groupID uuid.UUID) (rotations []Rotation, err error)
	RegisterClick(ctx context.Context, rotationID uuid.UUID) (err error)
	RegisterShown(ctx context.Context, rotationID uuid.UUID) (err error)
}
