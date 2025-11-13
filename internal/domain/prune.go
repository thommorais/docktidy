package domain

// PruneResult represents the result of a pruning operation
type PruneResult struct {
	ResourcesPruned int
	SpaceReclaimed  int64
	Errors          []error
}

// PruneOptions configures pruning behavior
type PruneOptions struct {
	DryRun        bool
	Force         bool
	OlderThanDays int
	IncludeTypes  []ResourceType
	ExcludeLabels []string
	MinSizeBytes  int64
}
