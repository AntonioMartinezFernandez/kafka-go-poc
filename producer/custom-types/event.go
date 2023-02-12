package custom_types

type Event struct {
	ID        int     `json:"id"`
	Device    string  `json:"device"`
	Timestamp string  `json:"timestamp"`
	Payload   Payload `json:"payload"`
}
