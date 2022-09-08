package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckRulesetParametersRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 任务ID
	TaskId string `json:"task_id"`

	// 规则集ID
	RulesetId string `json:"ruleset_id"`

	// 规则集语言
	Language string `json:"language"`

	// 分页索引，偏移量，非必填
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的数量，非必填
	Limit *int32 `json:"limit,omitempty"`
}

func (o CheckRulesetParametersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRulesetParametersRequest struct{}"
	}

	return strings.Join([]string{"CheckRulesetParametersRequest", string(data)}, " ")
}
