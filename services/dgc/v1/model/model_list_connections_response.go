package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListConnectionsResponse struct {
	Total          *string           `json:"total,omitempty"`
	Connections    *[]ConnectionInfo `json:"connections,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListConnectionsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ListConnectionsResponse", string(data)}, " ")
}
