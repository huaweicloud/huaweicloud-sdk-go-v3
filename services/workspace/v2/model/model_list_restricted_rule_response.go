package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRestrictedRuleResponse Response Object
type ListRestrictedRuleResponse struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`

	// 规则列表。
	Items          *[]AppRule `json:"items,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListRestrictedRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestrictedRuleResponse struct{}"
	}

	return strings.Join([]string{"ListRestrictedRuleResponse", string(data)}, " ")
}
