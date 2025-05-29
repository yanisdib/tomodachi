package hub

type CreateHubRequest struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description,omitempty"`
	IconURL     string   `json:"icon_url,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	EventType   string   `json:"event_type" validate:"required"`
	Languages   []string `json:"languages,omitempty"`
	Limit       int16    `json:"user_limit" validate:"required"`
	AccessRule  string   `json:"access_rule" validate:"required"`
	Password    string   `json:"password,omitempty"`
	Code        int8     `json:"code,omitempty"`
	StartAt     int64    `json:"start_at" validate:"required"`
	EndAt       int64    `json:"end_at,omitempty"`
	IsActive    bool     `json:"is_active" validate:"required"`
	IsOpen      bool     `json:"is_open" validate:"required"`
	ServerID    int64    `json:"server_id" validate:"required"`
	HostID      int64    `json:"host_id" validate:"required"`
}

type UpdateHubRequest struct {
	ID          int64    `json:"id" validate:"required"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	IconURL     string   `json:"icon_url,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	EventType   string   `json:"event_type,omitempty"`
	Languages   []string `json:"languages,omitempty"`
	Limit       int16    `json:"user_limit,omitempty"`
	AccessRule  string   `json:"access_rule,omitempty"`
	Password    string   `json:"password,omitempty"`
	Code        int8     `json:"code,omitempty"`
	StartAt     int64    `json:"start_at,omitempty"`
	EndAt       int64    `json:"end_at,omitempty"`
	IsActive    bool     `json:"is_active,omitempty"`
	IsOpen      bool     `json:"is_open,omitempty"`
	ServerID    int64    `json:"server_id,omitempty"`
	HostID      int64    `json:"host_id,omitempty"`
}

type DeleteHubRequest struct {
	ID int64 `json:"id" validate:"required"`
}
