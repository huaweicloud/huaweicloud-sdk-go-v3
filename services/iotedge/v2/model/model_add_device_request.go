package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AddDeviceRequest struct {
	// 边缘节点ID

	EdgeNodeId string `json:"edge_node_id"`

	Body *AddDeviceRequestBody `json:"body,omitempty"`
}

func (o AddDeviceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddDeviceRequest struct{}"
	}

	return strings.Join([]string{"AddDeviceRequest", string(data)}, " ")
}
