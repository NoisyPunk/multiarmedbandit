package queue

import (
	"time"
)

type Event struct {
	Name        string    `json:"name"`
	SlotID      string    `json:"slotId"`
	BannerID    string    `json:"bannerId"`
	GroupID     string    `json:"groupId"`
	DateAndTime time.Time `json:"dateAndTime"`
}
