package domain

import "time"

// UsageHistory tracks when a resource was last used
type UsageHistory struct {
	ResourceID   string
	ResourceType ResourceType
	LastAccessed time.Time
	AccessCount  int
}

// PruneHistory tracks pruning operations
type PruneHistory struct {
	ID        string
	Timestamp time.Time
	Result    PruneResult
	Options   PruneOptions
}
