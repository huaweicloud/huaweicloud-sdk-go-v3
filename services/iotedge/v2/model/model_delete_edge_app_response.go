package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteEdgeAppResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteEdgeAppResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteEdgeAppResponse struct{}"
	}

	return strings.Join([]string{"DeleteEdgeAppResponse", string(data)}, " ")
}
