package discord

type CreateDiscordServerRequest struct {
	Name             string `json:"name" validate:"required"`
	NitroLevel       int8   `json:"nitro_level" validate:"required"`
	IconURL          string `json:"icon_url,omitempty"`
	VideoQualityMode int8   `json:"video_quality_mode,omitempty"`
	AverageLatency   int64  `json:"average_latency,omitempty"`
	MemberCount      int64  `json:"member_count,omitempty"`
	SharedLink       string `json:"shared_link,omitempty"`
	Region           string `json:"region,omitempty"`
	HostID           int64  `json:"host_id" validate:"required"`
	CreatedAt        string `json:"created_at" validate:"required"`
}
