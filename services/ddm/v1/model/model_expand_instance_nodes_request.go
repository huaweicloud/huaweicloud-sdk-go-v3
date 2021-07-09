package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ExpandInstanceNodesRequest struct {
	// DDM实例ID

	InstanceId string `json:"instance_id"`

	Body *EnlargeRequest `json:"body,omitempty"`
}

func (o ExpandInstanceNodesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExpandInstanceNodesRequest struct{}"
	}

	return strings.Join([]string{"ExpandInstanceNodesRequest", string(data)}, " ")
}
