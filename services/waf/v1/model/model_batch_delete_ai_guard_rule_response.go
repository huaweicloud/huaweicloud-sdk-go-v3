package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAiGuardRuleResponse Response Object
type BatchDeleteAiGuardRuleResponse struct {

	// 已删除的规则ID列表
	RuleIds        *[]string `json:"rule_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteAiGuardRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAiGuardRuleResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteAiGuardRuleResponse", string(data)}, " ")
}
