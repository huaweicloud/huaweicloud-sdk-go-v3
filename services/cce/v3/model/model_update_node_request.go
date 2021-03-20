package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateNodeRequest struct {
	ClusterId string `json:"cluster_id"`

	NodeId string `json:"node_id"`

	ErrorStatus *string `json:"errorStatus,omitempty"`

	Body *CceClusterNodeInformation `json:"body,omitempty"`
}

func (o UpdateNodeRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateNodeRequest struct{}"
	}

	return strings.Join([]string{"UpdateNodeRequest", string(data)}, " ")
}
