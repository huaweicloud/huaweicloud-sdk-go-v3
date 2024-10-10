package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWafCustomRuleResponse Response Object
type ListWafCustomRuleResponse struct {

	// total
	Total *int32 `json:"total,omitempty"`

	// items
	Items          *[]WafCustomRule `json:"items,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListWafCustomRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWafCustomRuleResponse struct{}"
	}

	return strings.Join([]string{"ListWafCustomRuleResponse", string(data)}, " ")
}
