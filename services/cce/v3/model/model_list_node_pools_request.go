package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListNodePoolsRequest struct {
	ClusterId string `json:"cluster_id"`

	ErrorStatus *string `json:"errorStatus,omitempty"`

	ShowDefaultNodePool *string `json:"showDefaultNodePool,omitempty"`
}

func (o ListNodePoolsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListNodePoolsRequest struct{}"
	}

	return strings.Join([]string{"ListNodePoolsRequest", string(data)}, " ")
}
