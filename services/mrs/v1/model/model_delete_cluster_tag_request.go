package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteClusterTagRequest struct {
	ClusterId string `json:"cluster_id"`

	Key string `json:"key"`
}

func (o DeleteClusterTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteClusterTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteClusterTagRequest", string(data)}, " ")
}
