// Package domain contains core business entities and logic for docktidy.
package domain

import "time"

// ResourceType represents the type of Docker resource
type ResourceType string

const (
	// ResourceTypeContainer represents a Docker container
	ResourceTypeContainer ResourceType = "container"
	// ResourceTypeImage represents a Docker image
	ResourceTypeImage ResourceType = "image"
	// ResourceTypeVolume represents a Docker volume
	ResourceTypeVolume ResourceType = "volume"
	// ResourceTypeNetwork represents a Docker network
	ResourceTypeNetwork ResourceType = "network"
)

// Resource represents a Docker resource
type Resource struct {
	ID        string
	Type      ResourceType
	Name      string
	Size      int64
	CreatedAt time.Time
	LastUsed  time.Time
	InUse     bool
	Labels    map[string]string
	Tags      []string
}

// PruneCandidate represents a resource that can be pruned
type PruneCandidate struct {
	Resource
	Reason       string
	DaysSinceUse int
	RiskLevel    RiskLevel
}

// RiskLevel indicates the safety level of pruning a resource
type RiskLevel string

const (
	// RiskLevelSafe indicates the resource can be safely removed
	RiskLevelSafe RiskLevel = "safe"
	// RiskLevelMedium indicates the resource should be reviewed before removal
	RiskLevelMedium RiskLevel = "medium"
	// RiskLevelHigh indicates the resource is risky to remove
	RiskLevelHigh RiskLevel = "high"
)
