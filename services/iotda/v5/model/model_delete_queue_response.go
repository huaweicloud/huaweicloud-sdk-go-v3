package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteQueueResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteQueueResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteQueueResponse struct{}"
	}

	return strings.Join([]string{"DeleteQueueResponse", string(data)}, " ")
}
