package health

import "encoding/json"

type status string

const (
	up           status = "UP"
	down                = "DOWN"
	outOfService        = "OUT OF SERVICE"
	unknown             = "UNKNOWN"
)

// Health is a health status struct
type Health struct {
	status status
	info   map[string]interface{}
}

// MarshalJSON is a custom JSON marshaller
func (h Health) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{}

	for k, v := range h.info {
		data[k] = v
	}

	data["status"] = h.status

	return json.Marshal(data)
}

// NewHealth return a new Health with status Down
func NewHealth() Health {
	h := Health{
		info: make(map[string]interface{}),
	}

	h.Down()

	return h
}

// AddInfo adds a info value to the Info map
func (h *Health) AddInfo(key string, value interface{}) {
	if h.info == nil {
		h.info = make(map[string]interface{})
	}

	h.info[key] = value
}

// GetInfo returns a value from the info map
func (h Health) GetInfo(key string) interface{} {
	return h.info[key]
}

// IsUnknown returns true if Status is Unknown
func (h Health) IsUnknown() bool {
	return h.status == unknown
}

// IsUp returns true if Status is Up
func (h Health) IsUp() bool {
	return h.status == up
}

// IsDown returns true if Status is Down
func (h Health) IsDown() bool {
	return h.status == down
}

// IsOutOfService returns true if Status is IsOutOfService
func (h Health) IsOutOfService() bool {
	return h.status == outOfService
}

// Down set the status to Down
func (h *Health) Down() {
	h.status = down
}

// OutOfService set the status to OutOfService
func (h *Health) OutOfService() {
	h.status = outOfService
}

// Unknown set the status to Unknown
func (h *Health) Unknown() {
	h.status = unknown
}

// Up set the status to Up
func (h *Health) Up() {
	h.status = up
}
