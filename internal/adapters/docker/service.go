// Package docker provides the Docker adapter implementation.
package docker

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	dockerclient "github.com/docker/docker/client"
	"github.com/thommorais/docktidy/internal/domain"
	"github.com/thommorais/docktidy/internal/ports"
)

const healthCheckTimeout = 5 * time.Second

var _ ports.DockerService = (*Service)(nil)

// ErrNotImplemented is returned for Docker operations that are not yet supported.
var ErrNotImplemented = errors.New("docker adapter: not implemented")

// dockerAPIClient defines the subset of the Docker SDK client we need.
type dockerAPIClient interface {
	Ping(ctx context.Context) (types.Ping, error)
}

// Service implements the ports.DockerService interface using the Docker SDK.
type Service struct {
	client dockerAPIClient
}

// Option configures the Service.
type Option func(*Service)

// WithClient injects a custom Docker client (useful for testing).
func WithClient(cli dockerAPIClient) Option {
	return func(s *Service) {
		s.client = cli
	}
}

// NewService creates a Service with the default Docker client unless overridden via options.
func NewService(opts ...Option) (*Service, error) {
	svc := &Service{}

	for _, opt := range opts {
		opt(svc)
	}

	if svc.client == nil {
		cli, err := dockerclient.NewClientWithOpts(
			dockerclient.FromEnv,
			dockerclient.WithAPIVersionNegotiation(),
		)
		if err != nil {
			return nil, fmt.Errorf("create docker client: %w", err)
		}
		svc.client = cli
	}

	return svc, nil
}

// ListResources is not implemented yet.
func (s *Service) ListResources(context.Context, []domain.ResourceType) ([]domain.Resource, error) {
	return nil, ErrNotImplemented
}

// PruneResources is not implemented yet.
func (s *Service) PruneResources(context.Context, []domain.PruneCandidate, domain.PruneOptions) (domain.PruneResult, error) {
	return domain.PruneResult{}, ErrNotImplemented
}

// GetResourceDetails is not implemented yet.
func (s *Service) GetResourceDetails(context.Context, domain.ResourceType, string) (*domain.Resource, error) {
	return nil, ErrNotImplemented
}

// IsHealthy checks if the Docker daemon is reachable via Ping.
func (s *Service) IsHealthy(ctx context.Context) error {
	if s == nil || s.client == nil {
		return errors.New("docker adapter: client not initialized")
	}

	healthCtx, cancel := context.WithTimeout(ctx, healthCheckTimeout)
	defer cancel()

	if _, err := s.client.Ping(healthCtx); err != nil {
		return fmt.Errorf("docker ping failed: %w", err)
	}

	return nil
}
