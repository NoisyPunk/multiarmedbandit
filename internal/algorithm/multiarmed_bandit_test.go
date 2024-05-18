package algorithm

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"

	"github.com/NoisyPunk/multiarmedbandit/internal/storage"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func TestChooseBanner(t *testing.T) {
	id1, _ := uuid.Parse("936da01f-9abd-4d9d-80c4-02af85c295a7")
	id2, _ := uuid.Parse("874213e5-5c2b-43e3-9c7f-1a3b5c7d8e9f")
	id3, _ := uuid.Parse("2f6b4c1a-1a3b-4c5d-86f2-9b3c4d5e6f7g")
	id4, _ := uuid.Parse("e5f2a1c3-4b5d-9a8f-7g6h-5j4k3l2m1n")

	tests := []struct {
		name        string
		rotations   []storage.Rotation
		expected    storage.Rotation
		expectedErr error
	}{
		{
			name:        "empty input",
			rotations:   []storage.Rotation{},
			expectedErr: ErrEmptyRotationList,
		},
		{
			name: "single rotation",
			rotations: []storage.Rotation{
				{
					ID:       id1,
					BannerID: id2,
					GroupID:  id3,
					SlotID:   id4,
					Clicks:   10,
					Shows:    0,
				},
			},
			expected: storage.Rotation{
				ID:       id1,
				BannerID: id2,
				GroupID:  id3,
				SlotID:   id4,
				Clicks:   10,
				Shows:    0,
			},
		},
		{
			name: "multiple rotations",
			rotations: []storage.Rotation{
				{
					ID:       id1,
					BannerID: id2,
					GroupID:  id3,
					SlotID:   id4,
					Clicks:   10,
					Shows:    0,
				},
				{
					ID:       id2,
					BannerID: id3,
					GroupID:  id4,
					SlotID:   id1,
					Clicks:   20,
					Shows:    0,
				},
				{
					ID:       id3,
					BannerID: id4,
					GroupID:  id1,
					SlotID:   id2,
					Clicks:   30,
					Shows:    0,
				},
			},
			expected: storage.Rotation{
				ID:       id2,
				BannerID: id4,
				GroupID:  id1,
				SlotID:   id2,
				Clicks:   30,
				Shows:    0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rand.New(rand.NewSource(time.Now().UnixNano()))
			bestRotation, err := ChooseBanner(tt.rotations)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
			}
			require.Equal(t, bestRotation, tt.expected)
		})
	}
}

func TestChooseBannerProportions(t *testing.T) {
	id1, _ := uuid.Parse("936da01f-9abd-4d9d-80c4-02af85c295a7")
	id2, _ := uuid.Parse("874213e5-5c2b-43e3-9c7f-1a3b5c7d8e9f")

	tests := []struct {
		name         string
		rotations    []storage.Rotation
		showQuantity int
	}{
		{
			name: "test show count",
			rotations: []storage.Rotation{
				{
					ID:       id1,
					BannerID: id1,
					GroupID:  id1,
					SlotID:   id1,
					Clicks:   10,
					Shows:    0,
				},
				{
					ID:       id2,
					BannerID: id2,
					GroupID:  id2,
					SlotID:   id2,
					Clicks:   30,
					Shows:    0,
				},
			},
			showQuantity: 100,
		},
	}

	for _, tt := range tests {
		bbr := 0
		wbr := 0
		t.Run(tt.name, func(t *testing.T) {
			rand.New(rand.NewSource(time.Now().UnixNano()))
			for i := 0; i < tt.showQuantity; i++ {
				bestRotation, _ := ChooseBanner(tt.rotations)
				if bestRotation.ID == id2 {
					bbr++
					continue
				}
				wbr++
			}
			require.True(t, bbr > wbr)
		})
	}
}
