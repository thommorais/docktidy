package ports

import (
	"context"

	"github.com/thommorais/docktidy/internal/domain"
)

// DockerService defines operations for interacting with Docker
type DockerService interface {
	// ListResources retrieves all resources of specified types
	ListResources(ctx context.Context, types []domain.ResourceType) ([]domain.Resource, error)

	// PruneResources removes resources matching the criteria
	PruneResources(ctx context.Context, candidates []domain.PruneCandidate, opts domain.PruneOptions) (domain.PruneResult, error)

	// GetResourceDetails fetches detailed information about a specific resource
	GetResourceDetails(ctx context.Context, resourceType domain.ResourceType, id string) (*domain.Resource, error)

	// IsHealthy checks if Docker daemon is accessible
	IsHealthy(ctx context.Context) error
}
