package hub

import (
	"database/sql"
	"fmt"
)

type Hub struct {
	ID          int64         `json:"id" postgres:"id"`
	Name        string        `json:"name" postgres:"name"`
	Description string        `json:"description,omitempty" postgres:"description"`
	IconURL     string        `json:"icon_url,omitempty" postgres:"icon_url"`
	Tags        []string      `json:"tags,omitempty" postgres:"tags"`
	EventType   string        `json:"event_type" postgres:"event_type"`
	Languages   []string      `json:"languages,omitempty" postgres:"languages"`
	Limit       sql.NullInt16 `json:"user_limit" postgres:"user_limit"`
	AccessRule  HubAccessRule `json:"access_rule" postgres:"access_rule"`
	Password    string        `json:"password,omitempty" postgres:"password"`
	StartAt     sql.NullInt64 `json:"start_at" postgres:"start_at"`
	EndAt       sql.NullInt64 `json:"end_at,omitempty" postgres:"end_at"`
	IsActive    bool          `json:"is_active" postgres:"active"`
	IsOpen      bool          `json:"is_open" postgres:"open"`
	CreatedAt   int64         `json:"created_at" postgres:"created_at"`
	UpdatedAt   sql.NullInt64 `json:"updated_at,omitempty" postgres:"updated_at"`
	ClosedAt    sql.NullInt64 `json:"closed_at,omitempty" postgres:"closed_at,omitempty"`
	ServerID    sql.NullInt64 `json:"server_id" postgres:"server_id"`
	HostID      int64         `json:"host_id" postgres:"host_id"`
}

type HubAccessRule string

const (
	Public              HubAccessRule = "public"
	Private             HubAccessRule = "private"
	PrivateWithCode     HubAccessRule = "code"
	PrivateWithRequest  HubAccessRule = "request"
	PrivateWithPassword HubAccessRule = "password"
	PrivateWithInvite   HubAccessRule = "invite"
)

func CheckAccessRule(rule HubAccessRule) bool {
	switch rule {
	case Public, Private, PrivateWithCode, PrivateWithInvite, PrivateWithPassword, PrivateWithRequest:
		return true
	default:
		return false
	}
}

func ParseAccessRule(s string) (HubAccessRule, error) {
	rule := HubAccessRule(s)
	isValidRule := CheckAccessRule(rule)
	if !isValidRule {
		return "", fmt.Errorf("invalid hub access rule: %s", s)
	}

	return rule, nil
}
