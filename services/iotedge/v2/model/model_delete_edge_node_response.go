package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteEdgeNodeResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteEdgeNodeResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteEdgeNodeResponse struct{}"
	}

	return strings.Join([]string{"DeleteEdgeNodeResponse", string(data)}, " ")
}
