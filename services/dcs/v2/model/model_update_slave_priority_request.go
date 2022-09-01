package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateSlavePriorityRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 分片ID。
	GroupId string `json:"group_id" xml:"group_id"`

	// 节点ID。
	NodeId string `json:"node_id" xml:"node_id"`

	Body *PriorityBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateSlavePriorityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSlavePriorityRequest struct{}"
	}

	return strings.Join([]string{"UpdateSlavePriorityRequest", string(data)}, " ")
}
