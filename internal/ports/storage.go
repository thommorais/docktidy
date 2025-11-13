package ports

import (
	"context"

	"github.com/thommorais/docktidy/internal/domain"
)

// StorageService defines operations for persisting data
type StorageService interface {
	// SaveUsageHistory records resource usage
	SaveUsageHistory(ctx context.Context, history domain.UsageHistory) error

	// GetUsageHistory retrieves usage history for a resource
	GetUsageHistory(ctx context.Context, resourceID string) (*domain.UsageHistory, error)

	// SavePruneHistory records a pruning operation
	SavePruneHistory(ctx context.Context, history domain.PruneHistory) error

	// ListPruneHistory retrieves historical pruning operations
	ListPruneHistory(ctx context.Context, limit int) ([]domain.PruneHistory, error)

	// Close closes the storage connection
	Close() error
}
