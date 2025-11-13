package docker

import (
	"context"
	"errors"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/thommorais/docktidy/internal/domain"
)

type fakeClient struct {
	pingErr error
}

func (f fakeClient) Ping(context.Context) (types.Ping, error) {
	return types.Ping{}, f.pingErr
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
