package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LlmRuleInfoPromptDetectOpts 提示词验证
type LlmRuleInfoPromptDetectOpts struct {

	// 提示词索引
	Content *string `json:"content,omitempty"`

	// 提示词注入检测
	InjectEnable *bool `json:"inject_enable,omitempty"`

	// 提示词合规检测
	ComplianceEnable *bool `json:"compliance_enable,omitempty"`

	Action *LlmRuleInfoPromptDetectOptsAction `json:"action,omitempty"`
}

func (o LlmRuleInfoPromptDetectOpts) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LlmRuleInfoPromptDetectOpts struct{}"
	}

	return strings.Join([]string{"LlmRuleInfoPromptDetectOpts", string(data)}, " ")
}
