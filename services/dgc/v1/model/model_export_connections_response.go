package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ExportConnectionsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportConnectionsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExportConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ExportConnectionsResponse", string(data)}, " ")
}
