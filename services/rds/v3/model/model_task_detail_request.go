package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskDetailRequest 获取任务信息请求体。
type TaskDetailRequest struct {

	// 任务流ID
	WorkflowId string `json:"workflow_id"`

	// 任务名
	WorkflowName string `json:"workflow_name"`
}

func (o TaskDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskDetailRequest struct{}"
	}

	return strings.Join([]string{"TaskDetailRequest", string(data)}, " ")
}
