package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateBotMRuleStatusRequestBody 批量更新规则启用状态的请求体
type BatchUpdateBotMRuleStatusRequestBody struct {

	// 批量更新的规则ID
	RuleIds *[]int32 `json:"rule_ids,omitempty"`

	// 批量更新的启用状态
	Enable *bool `json:"enable,omitempty"`
}

func (o BatchUpdateBotMRuleStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateBotMRuleStatusRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateBotMRuleStatusRequestBody", string(data)}, " ")
}
