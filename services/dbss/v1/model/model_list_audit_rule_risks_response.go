package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditRuleRisksResponse Response Object
type ListAuditRuleRisksResponse struct {

	// 风险规则列表
	Rules *[]RuleRiskResponseRules `json:"rules,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuditRuleRisksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditRuleRisksResponse struct{}"
	}

	return strings.Join([]string{"ListAuditRuleRisksResponse", string(data)}, " ")
}
