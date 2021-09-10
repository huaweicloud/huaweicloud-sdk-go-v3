package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowClusterDetailRequest struct {
	// 集群ID

	ClusterId string `json:"cluster_id"`
}

func (o ShowClusterDetailRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowClusterDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowClusterDetailRequest", string(data)}, " ")
}
