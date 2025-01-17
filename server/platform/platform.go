package platform

// Platform defines the details of a hosting plaform
type Platform struct {
	ID              int               `pg:"id" json:"id"`
	Name            SupportedPlatform `pg:"name" json:"name"`
	Server          string            `pg:"server" json:"server,omitempty"`
	NitroLevel      uint8             `pg:"nitro_level" json:"nitroLevel,omitempty"`
	VideoResolution string            `pg:"video_resolution" json:"videoResolution,omitempty"`
	NetworkQuality  uint8             `pg:"network_quality" json:"networkQuality,omitempty"`
	AccessURL       string            `pg:"access_url" json:"accessURL,omitempty"`
}

// SupportedPlatform defines the platform used to host an event
type SupportedPlatform int

const (
	Discord SupportedPlatform = iota
	Zoom
	Teams
	Hangouts
)
