package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LlmRuleInfoRespDetectOpts 响应合规检测
type LlmRuleInfoRespDetectOpts struct {

	// 检测开关
	Enable *bool `json:"enable,omitempty"`

	// 响应内容索引
	ContentIdx *string `json:"content_idx,omitempty"`

	Action *LlmRuleInfoRespDetectOptsAction `json:"action,omitempty"`
}

func (o LlmRuleInfoRespDetectOpts) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LlmRuleInfoRespDetectOpts struct{}"
	}

	return strings.Join([]string{"LlmRuleInfoRespDetectOpts", string(data)}, " ")
}
