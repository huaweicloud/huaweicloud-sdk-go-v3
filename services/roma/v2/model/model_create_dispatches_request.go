package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDispatchesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 任务ID
	TaskId string `json:"task_id"`

	Body *TaskDispatch `json:"body,omitempty"`
}

func (o CreateDispatchesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDispatchesRequest struct{}"
	}

	return strings.Join([]string{"CreateDispatchesRequest", string(data)}, " ")
}
