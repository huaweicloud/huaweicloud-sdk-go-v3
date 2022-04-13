package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTaskRulesetRequest struct {
	// 项目ID

	ProjectId string `json:"project_id"`
	// 任务ID

	TaskId string `json:"task_id"`
}

func (o ListTaskRulesetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskRulesetRequest struct{}"
	}

	return strings.Join([]string{"ListTaskRulesetRequest", string(data)}, " ")
}
