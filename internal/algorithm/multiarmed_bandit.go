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
	epsilon := 0.2
	if rand.Float64() < epsilon { //nolint: gosec
		return randomRotation(rotations), err
	}
	bestRotation = rotations[0]
	return bestRotation, err
}

func randomRotation(rotations []storage.Rotation) (randomRotation storage.Rotation) {
	bannerIDs := make([]storage.Rotation, 0, len(rotations))
	bannerIDs = append(bannerIDs, rotations...)
	return bannerIDs[rand.Intn(len(bannerIDs))] //nolint: gosec
}
