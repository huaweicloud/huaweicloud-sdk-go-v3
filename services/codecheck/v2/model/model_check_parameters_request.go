package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckParametersRequest struct {
	// 项目ID

	ProjectId string `json:"project_id"`
	// 任务ID

	TaskId string `json:"task_id"`
	// 规则集ID

	RulesetId string `json:"ruleset_id"`
	// 规则集语言

	Language string `json:"language"`
}

func (o CheckParametersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckParametersRequest struct{}"
	}

	return strings.Join([]string{"CheckParametersRequest", string(data)}, " ")
}
