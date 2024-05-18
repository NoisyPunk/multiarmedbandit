package storage

import (
	"context"
	"fmt"

	"github.com/NoisyPunk/multiarmedbandit/internal/logger"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // for db
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	ErrDBConnection   = fmt.Errorf("cannot open pgx driver: ")
	ErrCreateBanner   = fmt.Errorf("cannot create banner: ")
	ErrCreateGroup    = fmt.Errorf("cannot create group: ")
	ErrCreateSlot     = fmt.Errorf("cannot create slot: ")
	ErrCreateRotation = fmt.Errorf("cannot create rotation: ")
)

type Storage struct {
	DB *sqlx.DB
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect(ctx context.Context, dsn string) (err error) {
	s.DB, err = sqlx.Open("postgres", dsn)
	if err != nil {
		return errors.Wrap(err, ErrDBConnection.Error())
	}
	return s.DB.PingContext(ctx)
}

func (s *Storage) Close() (err error) {
	return s.DB.Close()
}

func (s *Storage) AddBanner(ctx context.Context, description string) (id uuid.UUID, err error) {
	l := logger.FromContext(ctx)

	query := `INSERT INTO banners (id, description) 
				VALUES(:id, :description)`
	bannerID := uuid.New()
	banner := Banner{
		ID:          bannerID,
		Description: description,
	}

	_, err = s.DB.NamedQuery(query, banner)
	if err != nil {
		l.Error(err.Error(), zap.String("banner_id:", bannerID.String()))
		return uuid.Nil, errors.Wrap(err, ErrCreateBanner.Error())
	}
	l.Info("banner created:", zap.String("banner_id", bannerID.String()))
	return bannerID, nil
}

func (s *Storage) AddGroup(ctx context.Context, description string) (id uuid.UUID, err error) {
	l := logger.FromContext(ctx)

	query := `INSERT INTO groups (id, description) 
				VALUES(:id, :description)`
	groupID := uuid.New()
	group := Group{
		ID:          groupID,
		Description: description,
	}

	_, err = s.DB.NamedQuery(query, group)
	if err != nil {
		l.Error(err.Error(), zap.String("group_id:", groupID.String()))
		return uuid.Nil, errors.Wrap(err, ErrCreateGroup.Error())
	}
	l.Info("group created:", zap.String("group_id", groupID.String()))
	return groupID, nil
}

func (s *Storage) AddSlot(ctx context.Context, description string) (id uuid.UUID, err error) {
	l := logger.FromContext(ctx)

	query := `INSERT INTO slots (id, description) 
				VALUES(:id, :description)`
	slotID := uuid.New()
	group := Slot{
		ID:          slotID,
		Description: description,
	}

	_, err = s.DB.NamedQuery(query, group)
	if err != nil {
		l.Error(err.Error(), zap.String("slot_id:", slotID.String()))
		return uuid.Nil, errors.Wrap(err, ErrCreateSlot.Error())
	}
	l.Info("slot created:", zap.String("slot_id", slotID.String()))
	return slotID, nil
}

func (s *Storage) AddRotation(ctx context.Context, bannerID, slotID, groupID uuid.UUID) (id uuid.UUID, err error) {
	l := logger.FromContext(ctx)

	query := `INSERT INTO rotations (id, banner_id, group_id, slot_id, clicks, shows) 
				VALUES(:id, :banner_id, :group_id, :slot_id, :clicks, :shows)`
	rotationID := uuid.New()
	rotation := Rotation{
		ID:       rotationID,
		BannerID: bannerID,
		GroupID:  groupID,
		SlotID:   slotID,
		Clicks:   0,
		Shows:    0,
	}

	_, err = s.DB.NamedQuery(query, rotation)
	if err != nil {
		l.Error(err.Error(), zap.String("rotation_id:", rotationID.String()))
		return uuid.Nil, errors.Wrap(err, ErrCreateRotation.Error())
	}
	l.Info("rotation created:", zap.String("event_id", rotationID.String()))
	return rotationID, nil
}

func (s *Storage) GetSlotRotations(ctx context.Context, slotID, groupID uuid.UUID) (rotations []Rotation, err error) {
	l := logger.FromContext(ctx)

	query := `SELECT * FROM rotations where slot_id = $1 and group_id = $2`

	err = s.DB.Select(&rotations, query, slotID, groupID)
	if err != nil {
		return nil, err
	}
	l.Info("rotation list generated:", zap.String("slot_id", slotID.String()))
	return rotations, nil
}

func (s *Storage) RegisterClick(ctx context.Context, rotationID uuid.UUID) (err error) {
	l := logger.FromContext(ctx)

	query := `UPDATE rotations SET clicks = clicks + 1 WHERE id = $1`

	_, err = s.DB.Exec(query, rotationID)
	if err != nil {
		l.Error(err.Error(), zap.String("rotation_id:", rotationID.String()))
		return err
	}
	l.Info("rotation updated:", zap.String("rotation_id:", rotationID.String()))
	return nil
}

func (s *Storage) RegisterShown(ctx context.Context, rotationID uuid.UUID) (err error) {
	l := logger.FromContext(ctx)

	query := `UPDATE rotations SET shows = shows + 1 WHERE id = $1`

	_, err = s.DB.Exec(query, rotationID)
	if err != nil {
		l.Error(err.Error(), zap.String("rotation_id:", rotationID.String()))
		return err
	}
	l.Info("rotation updated:", zap.String("rotation_id:", rotationID.String()))
	return nil
}
