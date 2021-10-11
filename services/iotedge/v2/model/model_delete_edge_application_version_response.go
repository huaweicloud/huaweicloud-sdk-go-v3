package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteEdgeApplicationVersionResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteEdgeApplicationVersionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteEdgeApplicationVersionResponse struct{}"
	}

	return strings.Join([]string{"DeleteEdgeApplicationVersionResponse", string(data)}, " ")
}
