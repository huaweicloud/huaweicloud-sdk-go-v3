package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRuleGroupsResponse Response Object
type ListRuleGroupsResponse struct {

	// 规则组总数
	Total *int32 `json:"total,omitempty"`

	// 规则组列表
	Groups         *[]ResponseGroup `json:"groups,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListRuleGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRuleGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListRuleGroupsResponse", string(data)}, " ")
}
