package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ExportConnectionsRequest struct {
}

func (o ExportConnectionsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExportConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ExportConnectionsRequest", string(data)}, " ")
}
