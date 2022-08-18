package events

import (
	api "lolesports/domain/service"
	"time"
)

type AllEventData struct {
	Event *api.Event           `json:"event" bson:"event"`
	Games map[string]*GameData `json:"games" bson:"games,omitempty"`
}

type GameData struct {
	LastFrame       int                       `json:"last_frame" bson:"last_frame"`
	LastWindowFrame int                       `json:"last_window_frame" bson:"last_window_frame"`
	GameMetadata    *api.GameMetadata         `json:"game_metadata" bson:"game_metadata,omitempty"`
	PerksMetadata   map[int]*api.PerkMetadata `json:"perks_metadata" bson:"perks_metadata"`
	GameStartTime   *time.Time                `json:"game_start_time" bson:"game_start_time,omitempty"`
	LastTimeChecked *time.Time                `json:"last_time_checked" bson:"last_time_checked,omitempty"`
	PausedAt        *time.Time                `json:"paused_at" bson:"paused_at"`
	PausedDuration  time.Duration             `json:"paused_duration" bson:"paused_duration"`
	Fetched         bool                      `json:"fetched" bson:"fetched"`
}
