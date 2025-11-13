package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInteractionRuleGroupsSummaryResponse Response Object
type ListInteractionRuleGroupsSummaryResponse struct {

	// 互动规则总数。
	Count *int32 `json:"count,omitempty"`

	// 互动规则库列表。
	InteractionRuleGroups *[]InteractionRuleGroupSummary `json:"interaction_rule_groups,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListInteractionRuleGroupsSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInteractionRuleGroupsSummaryResponse struct{}"
	}

	return strings.Join([]string{"ListInteractionRuleGroupsSummaryResponse", string(data)}, " ")
}
