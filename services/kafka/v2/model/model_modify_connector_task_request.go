package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyConnectorTaskRequest Request Object
type ModifyConnectorTaskRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 任务ID。
	TaskId string `json:"task_id"`

	Body *SmartConnectTaskEntity `json:"body,omitempty"`
}

func (o ModifyConnectorTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyConnectorTaskRequest struct{}"
	}

	return strings.Join([]string{"ModifyConnectorTaskRequest", string(data)}, " ")
}
