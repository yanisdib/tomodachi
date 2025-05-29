package hub

import (
	"context"

	"github.com/yanisdib/tomodachi/domain/hub"
)

type HubService struct {
	repo hub.HubRepository
}

func NewHubService(r hub.HubRepository) *HubService {
	return &HubService{repo: r}
}

// CreateHub handles the creation of a new hub.
// It takes a CreateHubRequest as input and returns a CreateHubResponse or an error.
// The context parameter allows for cancellation and timeout control.
// It uses the repository to store the hub data in the database.
func (s *HubService) CreateHub(input *hub.CreateHubRequest, ctx context.Context) (*hub.CreateHubResponse, error) {
	return s.repo.Store(input, ctx)
}

// GetHubs retrieves a list of all hubs from the repository.
// It returns a slice of Hub pointers or an error if the operation fails.
// The context parameter allows for cancellation and timeout control.
// It uses the service to interact with the repository and fetch the hubs.
func (s *HubService) GetHubs(ctx context.Context) ([]*hub.GetAllHubsResponse, error) {
	return s.repo.GetHubs(ctx)
}
