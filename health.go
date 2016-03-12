package health

import "encoding/json"

type healthStatus int

const (
	UP healthStatus = iota
	DOWN
	OUT_OF_SERVICE
	UNKNOWN
)

func (h healthStatus) String() string {
	switch h {
	case UP:
		return "UP"
	case DOWN:
		return "DOWN"
	case OUT_OF_SERVICE:
		return "OUT_OF_SERVICE"
	case UNKNOWN:
		return "UNKNOWN"
	}

	return "NOT RECOGNIZED"
}

// Health is a health status struct
type Health struct {
	status healthStatus
	Info   map[string]interface{}
}

// MarshalJSON is a custom JSON marshaller
func (h Health) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{}
	data["status"] = h.status.String()

	for k, v := range h.Info {
		data[k] = v
	}

	return json.Marshal(data)
}

// NewHealth return a new Health with status Down
func NewHealth() Health {
	h := Health{
		Info: make(map[string]interface{}),
	}
	h.Down()

	return h
}

// IsUnknown returns true if Status is Unknown
func (h Health) IsUnknown() bool {
	return h.status == UNKNOWN
}

// IsUp returns true if Status is Up
func (h Health) IsUp() bool {
	return h.status == UP
}

// IsDown returns true if Status is Down
func (h Health) IsDown() bool {
	return h.status == DOWN
}

// IsOutOfService returns true if Status is IsOutOfService
func (h Health) IsOutOfService() bool {
	return h.status == OUT_OF_SERVICE
}

// Down set the status to Down
func (h *Health) Down() {
	h.status = DOWN
}

// OutOfService set the status to OutOfService
func (h *Health) OutOfService() {
	h.status = OUT_OF_SERVICE
}

// Unknown set the status to Unknown
func (h *Health) Unknown() {
	h.status = UNKNOWN
}

// Up set the status to Up
func (h *Health) Up() {
	h.status = UP
}
