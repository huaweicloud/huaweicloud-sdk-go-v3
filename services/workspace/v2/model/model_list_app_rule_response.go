package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppRuleResponse Response Object
type ListAppRuleResponse struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`

	// 规则列表。
	Items          *[]AppRule `json:"items,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListAppRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppRuleResponse struct{}"
	}

	return strings.Join([]string{"ListAppRuleResponse", string(data)}, " ")
}
