package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchStartOrStopTasksRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *OperateTasksBean `json:"body,omitempty" xml:"body"`
}

func (o BatchStartOrStopTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartOrStopTasksRequest struct{}"
	}

	return strings.Join([]string{"BatchStartOrStopTasksRequest", string(data)}, " ")
}
