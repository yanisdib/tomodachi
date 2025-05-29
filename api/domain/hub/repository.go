package hub

import (
	"context"
)

type HubRepository interface {
	Store(input *CreateHubRequest, ctx context.Context) (*CreateHubResponse, error)
	GetHubs(ctx context.Context) ([]*GetAllHubsResponse, error)
	// GetHubByID(id int64, ctx context.Context) (*Hub, error)
	// UpdateHub(updatedHub *Hub, ctx context.Context) (*Hub, error)
	// DeleteHub(id int64, ctx context.Context) error
}
