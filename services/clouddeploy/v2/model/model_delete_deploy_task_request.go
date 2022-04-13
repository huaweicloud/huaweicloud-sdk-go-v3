package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteDeployTaskRequest struct {
	// 部署任务id

	TaskId string `json:"task_id"`
}

func (o DeleteDeployTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeployTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteDeployTaskRequest", string(data)}, " ")
}
