package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTaskCmetricsRequest struct {
	// 项目ID

	ProjectId string `json:"project_id"`
	// 任务ID

	TaskId string `json:"task_id"`
}

func (o ShowTaskCmetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskCmetricsRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskCmetricsRequest", string(data)}, " ")
}
