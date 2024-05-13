package storage

import (
	"context"
	"github.com/google/uuid"
)

type Banner struct {
	Id          uuid.UUID `db:"id"`
	Description string    `db:"description"`
}

type Group struct {
	Id          uuid.UUID `db:"id"`
	Description string    `db:"description"`
}

type Slot struct {
	Id          uuid.UUID `db:"id"`
	Description string    `db:"description"`
	banners     []Banner  `db:"banners"`
}

type Rotation struct {
	Id       uuid.UUID `db:"id"`
	BannerId uuid.UUID `db:"banner_id"`
	GroupId  uuid.UUID `db:"group_id"`
	SlotId   uuid.UUID `db:"slot_id"`
	Clicks   int       `db:"clicks"`
	Shows    int       `db:"shows"`
}

type Storager interface {
	Connect(ctx context.Context, dsn string) (err error)
	Close() error
	AddBanner(ctx context.Context, description string) (id uuid.UUID, err error)

	AddGroup(ctx context.Context, description string) (id uuid.UUID, err error)

	AddSlot(ctx context.Context, description string) (id uuid.UUID, err error)

	AddRotation(ctx context.Context, bannerId, slotId, groupId uuid.UUID) (id uuid.UUID, err error)
	GetRotationsForSlot(ctx context.Context, slotId, groupId uuid.UUID) (rotations []Rotation, err error)
	RegisterClick(ctx context.Context, rotationId uuid.UUID) (err error)
	RegisterShown(ctx context.Context, rotationId uuid.UUID) (err error)
}
