package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchDeleteClusterTagsRequest struct {
	ClusterId string `json:"cluster_id"`

	Body *BatchDeleteClusterTagsReq `json:"body,omitempty"`
}

func (o BatchDeleteClusterTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteClusterTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteClusterTagsRequest", string(data)}, " ")
}
