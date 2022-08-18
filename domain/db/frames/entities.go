package frames

import (
	api "lolesports/domain/service"
)

type AllFrameData struct {
	EventID     string           `json:"event_id" bson:"event_id"`
	GameID      string           `json:"game_id" bson:"game_id"`
	Frame       *api.Frame       `json:"frame" bson:"frame,omitempty"`
	WindowFrame *api.WindowFrame `json:"window_frame" bson:"window_frame,omitempty"`
	Second      int              `json:"second" bson:"second"`
}
