package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteLogStreamResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteLogStreamResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteLogStreamResponse struct{}"
	}

	return strings.Join([]string{"DeleteLogStreamResponse", string(data)}, " ")
}
