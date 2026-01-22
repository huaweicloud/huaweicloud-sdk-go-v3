package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePullTaskRequest Request Object
type DeletePullTaskRequest struct {

	// 任务执行区域
	Region *string `json:"region,omitempty"`

	// 任务id
	TaskId string `json:"task_id"`
}

func (o DeletePullTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePullTaskRequest struct{}"
	}

	return strings.Join([]string{"DeletePullTaskRequest", string(data)}, " ")
}
