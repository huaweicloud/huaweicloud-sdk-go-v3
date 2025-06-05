package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTaskDefaultResultRequest Request Object
type CreateTaskDefaultResultRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 任务uri
	TaskUri string `json:"task_uri"`

	Body *InitExecuteTaskInfo `json:"body,omitempty"`
}

func (o CreateTaskDefaultResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTaskDefaultResultRequest struct{}"
	}

	return strings.Join([]string{"CreateTaskDefaultResultRequest", string(data)}, " ")
}
