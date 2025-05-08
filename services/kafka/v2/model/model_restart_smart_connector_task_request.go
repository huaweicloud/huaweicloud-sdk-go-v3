package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartSmartConnectorTaskRequest Request Object
type RestartSmartConnectorTaskRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// Smart Connect任务ID。
	TaskId string `json:"task_id"`
}

func (o RestartSmartConnectorTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartSmartConnectorTaskRequest struct{}"
	}

	return strings.Join([]string{"RestartSmartConnectorTaskRequest", string(data)}, " ")
}
