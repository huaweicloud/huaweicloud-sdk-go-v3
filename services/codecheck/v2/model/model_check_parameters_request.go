package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckParametersRequest struct {
	// 设置媒体类型和编码格式

	ContentType string `json:"Content-Type"`
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
