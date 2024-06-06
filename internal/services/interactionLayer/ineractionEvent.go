package interactionLayer

type EventType string

const (
	ADD_IMAGE   EventType = "ADD_IMAGE"
	MOVE_IMAGE  EventType = "MOVE_IMAGE"
	SCALE_IMAGE EventType = "SCALE_IMAGE"
	DEL_IMAGE   EventType = "DEL_IMAGE"
)

type InteractionEvent struct {
	EventType EventType `json:"eventType"`
	ID        string    `json:"Id"`
}
