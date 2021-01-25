package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateConnectionRequest struct {
	Body *ConnectionInfo `json:"body,omitempty"`
}

func (o CreateConnectionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateConnectionRequest struct{}"
	}

	return strings.Join([]string{"CreateConnectionRequest", string(data)}, " ")
}
