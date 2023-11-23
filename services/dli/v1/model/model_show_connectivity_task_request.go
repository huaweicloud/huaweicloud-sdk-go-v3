package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConnectivityTaskRequest Request Object
type ShowConnectivityTaskRequest struct {
	QueueName string `json:"queue_name"`

	TaskId string `json:"task_id"`
}

func (o ShowConnectivityTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConnectivityTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowConnectivityTaskRequest", string(data)}, " ")
}
