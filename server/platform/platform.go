package platform

// PlatformDetails defines the details of the hosting plaform
type Platform struct {
	Name       SupportedPlatform `json:"name"`
	Server     string            `json:"server,omitempty"`
	NitroLevel uint8             `json:"nitroLevel,omitempty"`
	Resolution string            `json:"resolution,omitempty"`
	Network    uint8             `json:"network,omitempty"`
	AccessLink string            `json:"accessLink,omitempty"`
}

// HostingPlatform defines the platform used to host an event
type SupportedPlatform int

const (
	Discord SupportedPlatform = iota
	Zoom
	Teams
	Hangouts
)
