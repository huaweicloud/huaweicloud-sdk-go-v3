package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListClusterTagsRequest struct {
	ClusterId string `json:"cluster_id"`
}

func (o ListClusterTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListClusterTagsRequest struct{}"
	}

	return strings.Join([]string{"ListClusterTagsRequest", string(data)}, " ")
}
