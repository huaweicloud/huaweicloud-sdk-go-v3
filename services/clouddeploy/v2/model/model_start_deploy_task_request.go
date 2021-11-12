package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StartDeployTaskRequest struct {
	// 部署任务id

	TaskId string `json:"task_id"`

	Body *EnvExecutionBody `json:"body,omitempty"`
}

func (o StartDeployTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartDeployTaskRequest struct{}"
	}

	return strings.Join([]string{"StartDeployTaskRequest", string(data)}, " ")
}
