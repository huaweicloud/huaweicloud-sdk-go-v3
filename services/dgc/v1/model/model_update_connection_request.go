package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateConnectionRequest struct {
	ConnectionName string `json:"connection_name"`

	Body *ConnectionInfo `json:"body,omitempty"`
}

func (o UpdateConnectionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateConnectionRequest struct{}"
	}

	return strings.Join([]string{"UpdateConnectionRequest", string(data)}, " ")
}
