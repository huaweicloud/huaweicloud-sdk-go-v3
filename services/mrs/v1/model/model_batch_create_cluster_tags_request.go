package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchCreateClusterTagsRequest struct {
	ClusterId string `json:"cluster_id"`

	Body *BatchCreateClusterTagsReq `json:"body,omitempty"`
}

func (o BatchCreateClusterTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateClusterTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateClusterTagsRequest", string(data)}, " ")
}
