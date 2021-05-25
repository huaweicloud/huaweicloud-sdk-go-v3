package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateSlavePriorityRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// 分片ID。

	GroupId string `json:"group_id"`
	// 节点ID。

	NodeId string `json:"node_id"`

	Body *PriorityBody `json:"body,omitempty"`
}

func (o UpdateSlavePriorityRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateSlavePriorityRequest struct{}"
	}

	return strings.Join([]string{"UpdateSlavePriorityRequest", string(data)}, " ")
}
