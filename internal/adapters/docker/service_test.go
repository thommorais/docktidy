package docker

import (
	"context"
	"errors"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/volume"
	"github.com/thommorais/docktidy/internal/domain"
)

type fakeClient struct {
	pingErr      error
	diskUsage    types.DiskUsage
	diskUsageErr error
}

func (f *fakeClient) Ping(context.Context) (types.Ping, error) {
	return types.Ping{}, f.pingErr
}

func (f *fakeClient) DiskUsage(context.Context, types.DiskUsageOptions) (types.DiskUsage, error) {
	if f.diskUsageErr != nil {
		return types.DiskUsage{}, f.diskUsageErr
	}
	return f.diskUsage, nil
}

func TestNewService_DefaultClient(t *testing.T) {
	svc, err := NewService()
	if err != nil {
		t.Fatalf("NewService() error = %v", err)
	}

	if svc.client == nil {
		t.Fatal("NewService() client is nil")
	}
}

func TestNewService_WithClient(t *testing.T) {
	custom := &fakeClient{}
	svc, err := NewService(WithClient(custom))
	if err != nil {
		t.Fatalf("NewService() error = %v", err)
	}

	if svc.client != custom {
		t.Fatal("NewService() did not use custom client")
	}
}

func TestService_IsHealthy(t *testing.T) {
	tests := []struct {
		name    string
		client  dockerAPIClient
		wantErr bool
	}{
		{
			name: "healthy",
			client: &fakeClient{
				pingErr: nil,
			},
			wantErr: false,
		},
		{
			name: "ping error",
			client: &fakeClient{
				pingErr: errors.New("socket unavailable"),
			},
			wantErr: true,
		},
		{
			name:    "missing client",
			client:  nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &Service{client: tt.client}
			err := svc.IsHealthy(context.Background())

			if tt.wantErr && err == nil {
				t.Fatalf("IsHealthy() error = nil, want error")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("IsHealthy() error = %v, want nil", err)
			}
		})
	}
}

func TestNotImplementedMethods(t *testing.T) {
	svc := &Service{}

	if _, err := svc.ListResources(context.Background(), nil); err != ErrNotImplemented {
		t.Errorf("ListResources() error = %v, want ErrNotImplemented", err)
	}

	if _, err := svc.GetResourceDetails(context.Background(), domain.ResourceTypeContainer, "id"); err != ErrNotImplemented {
		t.Errorf("GetResourceDetails() error = %v, want ErrNotImplemented", err)
	}

	if _, err := svc.PruneResources(context.Background(), nil, domain.PruneOptions{}); err != ErrNotImplemented {
		t.Errorf("PruneResources() error = %v, want ErrNotImplemented", err)
	}
}

func TestService_DiskUsage(t *testing.T) {
	cli := &fakeClient{
		diskUsage: types.DiskUsage{
			Images: []*image.Summary{
				{Size: 2000, Containers: 0, SharedSize: 150},
				{Size: 1000, Containers: 2, SharedSize: 50},
			},
			Containers: []*types.Container{
				{State: "running", SizeRw: 500},
				{State: "exited", SizeRw: 800},
			},
			Volumes: []*volume.Volume{
				{UsageData: &volume.UsageData{Size: 1024, RefCount: 1}},
				{UsageData: &volume.UsageData{Size: 2048, RefCount: 0}},
			},
			BuildCache: []*types.BuildCache{
				{InUse: true, Size: 300},
				{InUse: false, Size: 700},
			},
		},
	}

	svc := &Service{client: cli}
	usage, err := svc.DiskUsage(context.Background())
	if err != nil {
		t.Fatalf("DiskUsage() error = %v", err)
	}

	if len(usage.Rows) != 4 {
		t.Fatalf("DiskUsage() rows = %d, want 4", len(usage.Rows))
	}

	assertRow := func(t *testing.T, row domain.DiskUsageRow, rType string, total, active int, size, reclaim int64) {
		t.Helper()
		if row.Type != rType {
			t.Fatalf("row.Type = %s, want %s", row.Type, rType)
		}
		if row.Total != total || row.Active != active {
			t.Fatalf("%s counts = (%d,%d), want (%d,%d)", rType, row.Total, row.Active, total, active)
		}
		if row.SizeBytes != size || row.ReclaimableBytes != reclaim {
			t.Fatalf("%s sizes = (%d,%d), want (%d,%d)", rType, row.SizeBytes, row.ReclaimableBytes, size, reclaim)
		}
	}

	assertRow(t, usage.Rows[0], diskUsageTypeImages, 2, 1, 3000, 150)
	assertRow(t, usage.Rows[1], diskUsageTypeContainers, 2, 1, 1300, 800)
	assertRow(t, usage.Rows[2], diskUsageTypeVolumes, 2, 1, 3072, 2048)
	assertRow(t, usage.Rows[3], diskUsageTypeBuildCache, 2, 1, 1000, 700)
}
