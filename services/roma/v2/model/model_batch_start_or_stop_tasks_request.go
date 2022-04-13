package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchStartOrStopTasksRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *OperateTasksBean `json:"body,omitempty"`
}

func (o BatchStartOrStopTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartOrStopTasksRequest struct{}"
	}

	return strings.Join([]string{"BatchStartOrStopTasksRequest", string(data)}, " ")
}
