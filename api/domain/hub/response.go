package hub

type CreateHubResponse struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description,omitempty"`
	IconURL     string   `json:"icon_url,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	EventType   string   `json:"event_type" validate:"required"`
	Languages   []string `json:"languages,omitempty"`
	Limit       *int16   `json:"user_limit" validate:"required"`
	AccessRule  string   `json:"access_rule" validate:"required"`
	StartAt     *int64   `json:"start_at" validate:"required"`
	EndAt       *int64   `json:"end_at,omitempty"`
	IsActive    bool     `json:"is_active" validate:"required"`
	IsOpen      bool     `json:"is_open" validate:"required"`
	CreatedAt   int64    `json:"created_at" validate:"required"`
	ServerID    *int64   `json:"server_id" validate:"required"`
	HostID      int64    `json:"host_id" validate:"required"`
}

type GetAllHubsResponse struct {
	ID          int64    `json:"id" validate:"required"`
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description,omitempty"`
	IconURL     string   `json:"icon_url,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	EventType   string   `json:"event_type" validate:"required"`
	Languages   []string `json:"languages,omitempty"`
	Limit       *int16   `json:"user_limit" validate:"required"`
	AccessRule  string   `json:"access_rule" validate:"required"`
	StartAt     *int64   `json:"start_at" validate:"required"`
	EndAt       *int64   `json:"end_at,omitempty"`
	IsActive    bool     `json:"is_active" validate:"required"`
	IsOpen      bool     `json:"is_open" validate:"required"`
	CreatedAt   int64    `json:"created_at" validate:"required"`
	UpdatedAt   *int64   `json:"updated_at,omitempty"`
	ClosedAt    *int64   `json:"closed_at,omitempty"`
	ServerID    *int64   `json:"server_id" validate:"required"`
	HostID      int64    `json:"host_id" validate:"required"`
}
