package algorithm

import (
	"fmt"
	"math/rand"

	"github.com/NoisyPunk/multiarmedbandit/internal/storage"
)

var ErrEmptyRotationList = fmt.Errorf("rotation list is empty")

func ChooseBanner(rotations []storage.Rotation) (bestRotation storage.Rotation, err error) {
	if len(rotations) == 0 {
		return bestRotation, ErrEmptyRotationList
	}
	epsilon := 0.1
	if rand.Float64() < epsilon { //nolint: gosec
		return randomRotation(rotations), err
	}
	maxClicks := -1
	for _, rotation := range rotations {
		if rotation.Clicks > maxClicks {
			maxClicks = rotation.Clicks
			bestRotation = rotation
		}
	}
	return bestRotation, err
}

func randomRotation(rotations []storage.Rotation) (randomRotation storage.Rotation) {
	bannerIDs := make([]storage.Rotation, 0, len(rotations))
	for _, rotation := range rotations {
		bannerIDs = append(bannerIDs, rotation)
	}
	return bannerIDs[rand.Intn(len(bannerIDs))] //nolint: gosec
}
