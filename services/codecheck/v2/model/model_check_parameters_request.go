package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckParametersRequest struct {

	// 项目ID
	ProjectId string `json:"project_id" xml:"project_id"`

	// 任务ID
	TaskId string `json:"task_id" xml:"task_id"`

	// 规则集ID
	RulesetId string `json:"ruleset_id" xml:"ruleset_id"`

	// 规则集语言
	Language string `json:"language" xml:"language"`
}

func (o CheckParametersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckParametersRequest struct{}"
	}

	return strings.Join([]string{"CheckParametersRequest", string(data)}, " ")
}
