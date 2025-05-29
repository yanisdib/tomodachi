package hub

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yanisdib/tomodachi/domain/hub"
	"github.com/yanisdib/tomodachi/pkg/hash"
	"github.com/yanisdib/tomodachi/pkg/utils"
)

type pgRepository struct {
	db *pgxpool.Pool
}

func NewHubRepository(db *pgxpool.Pool) hub.HubRepository {
	return &pgRepository{db: db}
}

// CreateHub stores a entity creation request data to database.
func (r *pgRepository) Store(input *hub.CreateHubRequest, ctx context.Context) (*hub.CreateHubResponse, error) {
	var row hub.Hub
	hashedPassword, err := hash.HashPassword(input.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	query := CreateHubQuery
	createdAt := time.Now().Unix() // returns int64 seconds since epoch

	err = r.db.QueryRow(
		ctx,
		query,
		input.Name,
		input.Description,
		input.IconURL,
		input.Tags,
		input.EventType,
		input.Languages,
		input.Limit,
		input.AccessRule,
		hashedPassword,
		input.StartAt,
		input.EndAt,
		input.IsActive,
		input.IsOpen,
		createdAt,
		input.ServerID,
		input.HostID,
	).Scan(
		&row.ID,
		&row.Name,
		&row.Description,
		&row.IconURL,
		&row.Tags,
		&row.EventType,
		&row.Languages,
		&row.Limit,
		&row.AccessRule,
		&row.StartAt,
		&row.EndAt,
		&row.IsActive,
		&row.IsOpen,
		&row.CreatedAt,
		&row.ServerID,
		&row.HostID,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create entity: %w", err)
	}

	isValidRule := hub.CheckAccessRule(hub.HubAccessRule(input.AccessRule))
	if !isValidRule {
		input.AccessRule = string(hub.Public)
	}

	output := &hub.CreateHubResponse{
		Name:        row.Name,
		Description: row.Description,
		IconURL:     row.IconURL,
		Tags:        row.Tags,
		EventType:   row.EventType,
		Languages:   row.Languages,
		Limit:       utils.NullInt16Ptr(row.Limit),
		AccessRule:  string(row.AccessRule),
		StartAt:     utils.NullInt64Ptr(row.StartAt),
		EndAt:       utils.NullInt64Ptr(row.EndAt),
		IsActive:    row.IsActive,
		IsOpen:      row.IsOpen,
		CreatedAt:   row.CreatedAt,
		ServerID:    utils.NullInt64Ptr(row.ServerID),
		HostID:      row.HostID,
	}

	return output, nil

}

func (r *pgRepository) GetHubs(ctx context.Context) ([]*hub.GetAllHubsResponse, error) {
	query := GetHubsQuery
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while fetching hubs: %w", err)
	}

	defer rows.Close()

	var hubs []*hub.GetAllHubsResponse
	for rows.Next() {
		var currentHub hub.Hub
		err := rows.Scan(
			&currentHub.ID,
			&currentHub.Name,
			&currentHub.Description,
			&currentHub.IconURL,
			&currentHub.Tags,
			&currentHub.EventType,
			&currentHub.Languages,
			&currentHub.Limit,
			&currentHub.AccessRule,
			&currentHub.StartAt,
			&currentHub.EndAt,
			&currentHub.IsActive,
			&currentHub.IsOpen,
			&currentHub.CreatedAt,
			&currentHub.UpdatedAt,
			&currentHub.ClosedAt,
			&currentHub.ServerID,
			&currentHub.HostID,
		)

		output := &hub.GetAllHubsResponse{
			ID:          currentHub.ID,
			Name:        currentHub.Name,
			Description: currentHub.Description,
			IconURL:     currentHub.IconURL,
			Tags:        currentHub.Tags,
			EventType:   currentHub.EventType,
			Languages:   currentHub.Languages,
			Limit:       utils.NullInt16Ptr(currentHub.Limit),
			AccessRule:  string(currentHub.AccessRule),
			StartAt:     utils.NullInt64Ptr(currentHub.StartAt),
			EndAt:       utils.NullInt64Ptr(currentHub.EndAt),
			IsActive:    currentHub.IsActive,
			IsOpen:      currentHub.IsOpen,
			CreatedAt:   currentHub.CreatedAt,
			UpdatedAt:   utils.NullInt64Ptr(currentHub.UpdatedAt),
			ClosedAt:    utils.NullInt64Ptr(currentHub.ClosedAt),
			ServerID:    utils.NullInt64Ptr(currentHub.ServerID),
			HostID:      currentHub.HostID,
		}

		if err != nil {
			return nil, fmt.Errorf("an error occurred while scanning hub row: %w", err)
		}

		hubs = append(hubs, output)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("an error occurred while iterating over hub rows: %w", err)
	}

	if len(hubs) == 0 {
		return nil, nil
	}

	return hubs, nil
}
