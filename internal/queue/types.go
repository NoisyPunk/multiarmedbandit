package queue

import (
	"time"
)

type Event struct {
	Name        string    `json:"name"`
	SlotID      string    `json:"slot_id"`
	BannerID    string    `json:"banner_id"`
	GroupID     string    `json:"group_id"`
	DateAndTime time.Time `json:"date_and_time"`
}
