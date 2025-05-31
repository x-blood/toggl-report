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

// FilterByWorkspaceID は指定されたワークスペースIDに一致するTimeEntryのみを返し、一致しないものを除外します
func FilterByWorkspaceID(entries []TimeEntry, workspaceID uint64) []TimeEntry {
	var filtered []TimeEntry
	for _, entry := range entries {
		if entry.WorkspaceID == workspaceID {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}
