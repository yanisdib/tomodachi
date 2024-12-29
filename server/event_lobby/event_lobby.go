package eventlobby

import (
	"yanisdib/tomodachi/platform"
	"yanisdib/tomodachi/user"
)

// Group defines a group of users for a streaming session
type EventLobby struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Host        *user.User         `json:"host"`
	Status      EventStatus        `json:"status"`
	Platform    *platform.Platform `json:"platform"`
	StartAt     uint64             `json:"startAt"`
	EndAt       uint64             `json:"endAt"`
	AccessRule  AccessRule         `json:"accessRule"`
	Limit       uint8              `json:"limit,omitempty"`
	CreatedAt   uint64             `json:"createdAt"`
	UpdatedAt   uint64             `json:"updatedAt,omitempty"`
}

// AccessRule is an enumeration that defines the access required to join a lobby
type AccessRule int

const (
	Public AccessRule = iota
	PrivateWithRequest
	PrivateWithPassword
	PrivateWithLink
)

// EventStatus is an enumeration that defines the status of an event
type EventStatus int

const (
	OnGoing EventStatus = iota
	Pending
	OnHold
	Closed
)
