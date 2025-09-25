package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LlmRuleInfo struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 规则名
	Name *string `json:"name,omitempty"`

	// 规则描述
	Discription *string `json:"discription,omitempty"`

	// 模型问答URL
	Url *string `json:"url,omitempty"`

	PromptDetectOpts *LlmRuleInfoPromptDetectOpts `json:"prompt_detect_opts,omitempty"`

	RespDetectOpts *LlmRuleInfoRespDetectOpts `json:"resp_detect_opts,omitempty"`
}

func (o LlmRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LlmRuleInfo struct{}"
	}

	return strings.Join([]string{"LlmRuleInfo", string(data)}, " ")
}
