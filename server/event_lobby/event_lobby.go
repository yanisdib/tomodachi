package eventlobby

import "yanisdib/tomodachi/user"

// Group defines a group of users for a streaming session
type EventLobby struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Host        *user.User      `json:"host"`
	Status      EventStatus     `json:"status"`
	Platform    PlatformDetails `json:"platform"`
	StartAt     uint64          `json:"startAt"`
	EndAt       uint64          `json:"endAt"`
	Mode        AccessRule      `json:"mode"`
	Limit       uint8           `json:"limit,omitempty"`
	CreatedAt   uint64          `json:"createdAt"`
	UpdatedAt   uint64          `json:"updatedAt,omitempty"`
}

// AccessRule is an enumeration that defines the access required to join a lobby
type AccessRule int

const (
	Public AccessRule = iota
	PrivateWithRequest
	PrivateWithPassword
	PrivateWithLink
)

// EventStatus defines the status of an event
type EventStatus int

const (
	OnGoing EventStatus = iota
	Pending
	OnHold
	Closed
)

// HostingPlatform defines the platform used to host an event
type HostingPlatform int

const (
	Discord HostingPlatform = iota
	Zoom
	Teams
	Hangouts
)

type PlatformDetails struct {
	Name       HostingPlatform `json:"name"`
	Server     string          `json:"server,omitempty"`
	NitroLevel uint8           `json:"nitroLevel,omitempty"`
	Resolution string          `json:"resolution,omitempty"`
	Network    uint8           `json:"network,omitempty"`
	AccessLink string          `json:"accessLink,omitempty"`
}
