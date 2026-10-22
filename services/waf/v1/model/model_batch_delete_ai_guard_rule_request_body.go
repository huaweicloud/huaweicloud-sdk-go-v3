package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAiGuardRuleRequestBody 批量删除规则请求
type BatchDeleteAiGuardRuleRequestBody struct {

	// 规则ID列表
	RuleIds []string `json:"rule_ids"`
}

func (o BatchDeleteAiGuardRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAiGuardRuleRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteAiGuardRuleRequestBody", string(data)}, " ")
}
