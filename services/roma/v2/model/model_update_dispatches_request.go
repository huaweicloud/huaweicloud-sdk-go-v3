package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateDispatchesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 任务ID
	TaskId string `json:"task_id" xml:"task_id"`

	// 调度唯一标识，调度计划ID
	DispatchId string `json:"dispatch_id" xml:"dispatch_id"`

	Body *TaskDispatch `json:"body,omitempty" xml:"body"`
}

func (o UpdateDispatchesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDispatchesRequest struct{}"
	}

	return strings.Join([]string{"UpdateDispatchesRequest", string(data)}, " ")
}
