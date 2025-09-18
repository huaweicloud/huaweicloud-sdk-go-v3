package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskDetailRequest Request Object
type ShowTaskDetailRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *TaskDetailRequest `json:"body,omitempty"`
}

func (o ShowTaskDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskDetailRequest", string(data)}, " ")
}
