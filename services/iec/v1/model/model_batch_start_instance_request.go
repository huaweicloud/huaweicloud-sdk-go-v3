package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchStartInstanceRequest struct {
	Body *BatchStartInstanceRequestBody `json:"body,omitempty"`
}

func (o BatchStartInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchStartInstanceRequest struct{}"
	}

	return strings.Join([]string{"BatchStartInstanceRequest", string(data)}, " ")
}
