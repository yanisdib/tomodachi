package platform

// Platform defines the details of a hosting plaform
type Platform struct {
	Id              int               `json:"id"`
	Name            SupportedPlatform `json:"name"`
	Server          string            `json:"server,omitempty"`
	NitroLevel      uint8             `json:"nitroLevel,omitempty"`
	VideoResolution string            `json:"videoResolution,omitempty"`
	NetworkQuality  uint8             `json:"networkQuality,omitempty"`
	AccessLink      string            `json:"accessLink,omitempty"`
}

// SupportedPlatform defines the platform used to host an event
type SupportedPlatform int

const (
	Discord SupportedPlatform = iota
	Zoom
	Teams
	Hangouts
)
