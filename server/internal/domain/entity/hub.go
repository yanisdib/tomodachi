package entity

type Hub struct {
	ID          int64         `json:"id" postgres:"id"`
	Name        string        `json:"name" postgres:"name"`
	Description string        `json:"description,omitempty" postgres:"description"`
	IconURL     string        `json:"icon_url,omitempty" postgres:"icon_url"`
	Tags        string        `json:"tags,omitempty" postgres:"tags"`
	EventType   string        `json:"event_type" postgres:"event_type"`
	Languages   []string      `json:"languages" postgres:"languages"`
	Limit       int8          `json:"limit" postgres:"limit"`
	AccessRule  HubAccessRule `json:"access_rule" postgres:"access_rule"`
	Password    string        `json:"password,omitempty" postgres:"password"`
	StartedAt   string        `json:"started_at" postgres:"started_at"`
	EndedAt     string        `json:"ended_at,omitempty" postgres:"ended_at"`
	IsActive    bool          `json:"is_active" postgres:"active"`
	IsOpen      bool          `json:"is_open" postgres:"open"`
	CreatedAt   string        `json:"created_at" postgres:"created_at"`
	UpdatedAt   string        `json:"updated_at,omitempty" postgres:"updated_at"`
	DeletedAt   string        `json:"deleted_at,omitempty" postgres:"deleted_at"`
	ServerID    int8          `json:"server_id" postgres:"server_id"`
	HostID      int64         `json:"host_id" postgres:"host_id"`
}

type HubAccessRule string

const (
	Public   HubAccessRule = "public"
	Private  HubAccessRule = "private"
	Password HubAccessRule = "password"
	Invite   HubAccessRule = "invite"
)
