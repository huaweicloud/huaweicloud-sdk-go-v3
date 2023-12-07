package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityDataClassificationRuleGroupsResponse Response Object
type ListSecurityDataClassificationRuleGroupsResponse struct {

	// 规则组列表
	RuleGroups *[]DataClassificationGroupQueryDto `json:"rule_groups,omitempty"`

	// 规则组总数
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSecurityDataClassificationRuleGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityDataClassificationRuleGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityDataClassificationRuleGroupsResponse", string(data)}, " ")
}
