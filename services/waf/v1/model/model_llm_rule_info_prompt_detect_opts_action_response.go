package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LlmRuleInfoPromptDetectOptsActionResponse 响应返回
type LlmRuleInfoPromptDetectOptsActionResponse struct {

	// 响应码
	StatusCode *int32 `json:"status_code,omitempty"`

	// Content_type
	ContentType *string `json:"content_type,omitempty"`

	// 响应页面内容
	Content *string `json:"content,omitempty"`
}

func (o LlmRuleInfoPromptDetectOptsActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LlmRuleInfoPromptDetectOptsActionResponse struct{}"
	}

	return strings.Join([]string{"LlmRuleInfoPromptDetectOptsActionResponse", string(data)}, " ")
}
