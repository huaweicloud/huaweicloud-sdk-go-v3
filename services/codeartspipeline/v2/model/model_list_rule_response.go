package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRuleResponse Response Object
type ListRuleResponse struct {

	// **参数解释**： 静态规则列表。 **取值范围**： 不涉及。
	Data *[]Rule `json:"data,omitempty"`

	// **参数解释**： 本次查询的静态规则总数。 **取值范围**： 不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRuleResponse struct{}"
	}

	return strings.Join([]string{"ListRuleResponse", string(data)}, " ")
}
