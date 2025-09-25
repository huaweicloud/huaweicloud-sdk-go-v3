package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LlmRuleInfoPromptDetectOptsAction 防护
type LlmRuleInfoPromptDetectOptsAction struct {

	// 防护动作
	Category *string `json:"category,omitempty"`

	Response *LlmRuleInfoPromptDetectOptsActionResponse `json:"response,omitempty"`
}

func (o LlmRuleInfoPromptDetectOptsAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LlmRuleInfoPromptDetectOptsAction struct{}"
	}

	return strings.Join([]string{"LlmRuleInfoPromptDetectOptsAction", string(data)}, " ")
}
