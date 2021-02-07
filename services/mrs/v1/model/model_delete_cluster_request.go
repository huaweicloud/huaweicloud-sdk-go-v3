package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteClusterRequest struct {
	ClusterId string `json:"cluster_id"`
}

func (o DeleteClusterRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteClusterRequest struct{}"
	}

	return strings.Join([]string{"DeleteClusterRequest", string(data)}, " ")
}
