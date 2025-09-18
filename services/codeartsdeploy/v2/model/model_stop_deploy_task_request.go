package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopDeployTaskRequest Request Object
type StopDeployTaskRequest struct {

	// task_id
	TaskId string `json:"task_id"`

	// record_id
	RecordId string `json:"record_id"`

	Body *string `json:"body,omitempty"`
}

func (o StopDeployTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopDeployTaskRequest struct{}"
	}

	return strings.Join([]string{"StopDeployTaskRequest", string(data)}, " ")
}
