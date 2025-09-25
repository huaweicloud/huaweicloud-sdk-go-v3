package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LlmRuleInfoRespDetectOptsAction 响应防护
type LlmRuleInfoRespDetectOptsAction struct {

	// 响应防护动作
	Category *string `json:"category,omitempty"`

	// 终止响应协议
	AbortResponseContent *string `json:"abort_response_content,omitempty"`
}

func (o LlmRuleInfoRespDetectOptsAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LlmRuleInfoRespDetectOptsAction struct{}"
	}

	return strings.Join([]string{"LlmRuleInfoRespDetectOptsAction", string(data)}, " ")
}
