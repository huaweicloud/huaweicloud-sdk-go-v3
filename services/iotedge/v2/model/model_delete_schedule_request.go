package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScheduleRequest Request Object
type DeleteScheduleRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 调度计划id
	ScheduleId string `json:"schedule_id"`
}

func (o DeleteScheduleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScheduleRequest struct{}"
	}

	return strings.Join([]string{"DeleteScheduleRequest", string(data)}, " ")
}
