package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScheduleRequest Request Object
type UpdateScheduleRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 调度计划id
	ScheduleId string `json:"schedule_id"`

	Body *UpdateScheduleReqDto `json:"body,omitempty"`
}

func (o UpdateScheduleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScheduleRequest struct{}"
	}

	return strings.Join([]string{"UpdateScheduleRequest", string(data)}, " ")
}
