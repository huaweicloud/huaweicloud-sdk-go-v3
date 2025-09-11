package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditRuleRisksNewResponse Response Object
type ListAuditRuleRisksNewResponse struct {

	// 风险规则列表
	Rules *[]RuleRiskResponseRules `json:"rules,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 自定义规则总数
	CustomizeTotal *int32 `json:"customize_total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuditRuleRisksNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditRuleRisksNewResponse struct{}"
	}

	return strings.Join([]string{"ListAuditRuleRisksNewResponse", string(data)}, " ")
}
