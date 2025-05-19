package entity

type DiscordServer struct {
	ID               int64  `json:"id" postgres:"id"`
	Name             string `json:"name" postgres:"name"`
	NitroLevel       int8   `json:"nitro_level" postgres:"nitro_level"`
	IconURL          string `json:"icon_url" postgres:"icon_url"`
	VideoQualityMode int8   `json:"video_quality_mode" postgres:"video_quality_mode"`
	AverageLatency   int64  `json:"average_latency" postgres:"average_latency"`
	MemberCount      int64  `json:"member_count" postgres:"member_count"`
	SharedLink       string `json:"shared_link" postgres:"shared_link"`
	Region           string `json:"region" postgres:"region"`
	HostID           int64  `json:"host_id" postgres:"host_id"`
	CreatedAt        string `json:"created_at" postgres:"created_at"`
	UpdatedAt        string `json:"updated_at,omitempty" postgres:"updated_at"`
	DeletedAt        string `json:"deleted_at,omitempty" postgres:"deleted_at"`
}
