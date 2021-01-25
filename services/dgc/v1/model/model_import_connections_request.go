package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ImportConnectionsRequest struct {
	Body *ImportConnectionReq `json:"body,omitempty"`
}

func (o ImportConnectionsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImportConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ImportConnectionsRequest", string(data)}, " ")
}
