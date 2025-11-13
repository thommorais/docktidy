package domain

import "time"

// ResourceType represents the type of Docker resource
type ResourceType string

const (
	ResourceTypeContainer ResourceType = "container"
	ResourceTypeImage     ResourceType = "image"
	ResourceTypeVolume    ResourceType = "volume"
	ResourceTypeNetwork   ResourceType = "network"
)

// Resource represents a Docker resource
type Resource struct {
	ID          string
	Type        ResourceType
	Name        string
	Size        int64
	CreatedAt   time.Time
	LastUsed    time.Time
	InUse       bool
	Labels      map[string]string
	Tags        []string
}

// PruneCandidate represents a resource that can be pruned
type PruneCandidate struct {
	Resource
	Reason      string
	DaysSinceUse int
	RiskLevel   RiskLevel
}

// RiskLevel indicates the safety level of pruning a resource
type RiskLevel string

const (
	RiskLevelSafe   RiskLevel = "safe"
	RiskLevelMedium RiskLevel = "medium"
	RiskLevelHigh   RiskLevel = "high"
)
