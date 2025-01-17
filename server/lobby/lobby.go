package lobby

import (
	"github.com/yanisdib/tomodachi/server/platform"
	"github.com/yanisdib/tomodachi/server/user"
)

// Group defines a group of users for a streaming session
type Lobby struct {
	ID          string             `pg:"id" json:"id"`
	Name        string             `pg:"name" json:"name"`
	Description string             `pg:"description" json:"description"`
	Host        *user.User         `pg:"host_id" json:"host"`
	Status      EventStatus        `pg:"status" json:"status"`
	Platform    *platform.Platform `pg:"platform_id" json:"platform"`
	StartAt     uint64             `pg:"start_at" json:"startAt"`
	EndAt       uint64             `pg:"end_at" json:"endAt"`
	AccessRule  AccessRule         `pg:"access_rule" json:"accessRule"`
	Limit       uint8              `pg:"users_limit" json:"limit,omitempty"`
	CreatedAt   uint64             `pg:"created_at" json:"createdAt"`
	UpdatedAt   uint64             `pg:"updated_at" json:"updatedAt,omitempty"`
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
