package platform

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PlatformManager interface {
	CreatePlatform(ctx context.Context, input Platform) (*Platform, error)
	UpdatePlatform(ctx context.Context, input Platform) error
	DeletePlatform(ctx context.Context, id int) error
}

// Repository defines a PostgreSQL repository
type Repository struct {
	db *pgxpool.Pool
}

// NewRepository creates  a new PostgreSQL repository
func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreatePlatform(ctx context.Context, input Platform) (*Platform, error) {
	var id int
	const query string = `INSERT INTO platforms(
		name, 
		server, 
		nitro_level, 
		video_resolution, 
		network_quality, 
		access_url
	) VALUES($1, $2, $3, $4, $5, $6) 
	RETURNING id`

	err := r.db.QueryRow(
		ctx,
		query,
		input.Name,
		input.Server,
		input.NitroLevel,
		input.VideoResolution,
		input.NetworkQuality,
		input.AccessURL,
	).Scan(&id)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, ErrDuplicate
			}
		}
	}

	input.ID = id

	return &input, nil
}
