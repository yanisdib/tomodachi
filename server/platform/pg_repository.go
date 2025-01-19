package platform

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// pgRepository defines a PostgreSQL repository
type pgRepository struct {
	db *pgxpool.Pool
}

// NewpgRepository creates  a new PostgreSQL pgRepository
func NewPgRepository(db *pgxpool.Pool) *pgRepository {
	return &pgRepository{
		db: db,
	}
}

// Store stores a new platform in the database and returns the created platform
func (r *pgRepository) Store(ctx context.Context, input *CreatePlatformInput) (*Platform, error) {
	var output Platform
	query := createPlatformQuery

	err := r.db.QueryRow(
		ctx,
		query,
		input.Name,
		input.Server,
		input.NitroLevel,
		input.VideoResolution,
		input.NetworkQuality,
		input.AccessURL,
	).Scan(&output)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, ErrDuplicate
			}
		}
	}

	return &output, nil
}
