package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInteractionRuleGroupsResponse Response Object
type ListInteractionRuleGroupsResponse struct {

	// **参数解释**： 互动规则总数。
	Count *int32 `json:"count,omitempty"`

	// 互动规则库列表。
	InteractionRuleGroups *[]InteractionRuleGroupDetail `json:"interaction_rule_groups,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListInteractionRuleGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInteractionRuleGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListInteractionRuleGroupsResponse", string(data)}, " ")
}
