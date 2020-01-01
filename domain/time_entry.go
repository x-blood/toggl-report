package domain

import "time"

type TimeEntry struct {
	ID          uint64    `json:"id"`
	GUID        string    `json:"guid"`
	WorkspaceID uint64    `json:"wid"`
	ProjectID   uint64    `json:"pid"`
	Start       time.Time `json:"start"`
	Stop        time.Time `json:"stop"`
	Duration    uint64    `json:"duration"`
	Description string    `json:"description"`
}
