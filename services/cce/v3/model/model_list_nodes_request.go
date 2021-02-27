package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListNodesRequest struct {
	ClusterId   string  `json:"cluster_id"`
	ErrorStatus *string `json:"errorStatus,omitempty"`
}

func (o ListNodesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListNodesRequest struct{}"
	}

	return strings.Join([]string{"ListNodesRequest", string(data)}, " ")
}
