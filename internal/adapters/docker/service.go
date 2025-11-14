// Package docker provides the Docker adapter implementation.
package docker

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/volume"
	dockerclient "github.com/docker/docker/client"
	"github.com/thommorais/docktidy/internal/domain"
	"github.com/thommorais/docktidy/internal/ports"
)

const healthCheckTimeout = 5 * time.Second
const diskUsageTypeImages = "Images"
const diskUsageTypeContainers = "Containers"
const diskUsageTypeVolumes = "Local Volumes"
const diskUsageTypeBuildCache = "Build Cache"

var _ ports.DockerService = (*Service)(nil)

// ErrNotImplemented is returned for Docker operations that are not yet supported.
var ErrNotImplemented = errors.New("docker adapter: not implemented")

// dockerAPIClient defines the subset of the Docker SDK client we need.
type dockerAPIClient interface {
	Ping(ctx context.Context) (types.Ping, error)
	DiskUsage(ctx context.Context, options types.DiskUsageOptions) (types.DiskUsage, error)
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

// DiskUsage retrieves disk usage information from Docker.
func (s *Service) DiskUsage(ctx context.Context) (domain.DiskUsage, error) {
	if s == nil || s.client == nil {
		return domain.DiskUsage{}, errors.New("docker adapter: client not initialized")
	}

	usageCtx, cancel := context.WithTimeout(ctx, healthCheckTimeout)
	defer cancel()

	data, err := s.client.DiskUsage(usageCtx, types.DiskUsageOptions{})
	if err != nil {
		return domain.DiskUsage{}, fmt.Errorf("docker disk usage: %w", err)
	}

	rows := []domain.DiskUsageRow{
		aggregateImageUsage(data.Images),
		aggregateContainerUsage(data.Containers),
		aggregateVolumeUsage(data.Volumes),
		aggregateBuildCacheUsage(data.BuildCache),
	}

	return domain.DiskUsage{Rows: rows}, nil
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

func aggregateImageUsage(images []*image.Summary) domain.DiskUsageRow {
	row := domain.DiskUsageRow{Type: diskUsageTypeImages}
	for _, img := range images {
		if img == nil {
			continue
		}
		row.Total++
		if img.Containers > 0 {
			row.Active++
		}
		row.SizeBytes += img.Size
		if img.Containers == 0 {
			row.ReclaimableBytes += img.SharedSize
		}
	}
	return row
}

func aggregateContainerUsage(containers []*types.Container) domain.DiskUsageRow {
	row := domain.DiskUsageRow{Type: diskUsageTypeContainers}
	for _, c := range containers {
		if c == nil {
			continue
		}
		row.Total++
		if strings.EqualFold(c.State, "running") {
			row.Active++
		}
		row.SizeBytes += c.SizeRw
		if !strings.EqualFold(c.State, "running") {
			row.ReclaimableBytes += c.SizeRw
		}
	}
	return row
}

func aggregateVolumeUsage(volumes []*volume.Volume) domain.DiskUsageRow {
	row := domain.DiskUsageRow{Type: diskUsageTypeVolumes}
	for _, v := range volumes {
		if v == nil {
			continue
		}
		row.Total++
		usage := v.UsageData
		if usage != nil {
			row.SizeBytes += usage.Size
			if usage.RefCount > 0 {
				row.Active++
			} else {
				row.ReclaimableBytes += usage.Size
			}
		}
	}
	return row
}

func aggregateBuildCacheUsage(cache []*types.BuildCache) domain.DiskUsageRow {
	row := domain.DiskUsageRow{Type: diskUsageTypeBuildCache}
	for _, c := range cache {
		if c == nil {
			continue
		}
		row.Total++
		if c.InUse {
			row.Active++
		} else {
			row.ReclaimableBytes += c.Size
		}
		row.SizeBytes += c.Size
	}
	return row
}
