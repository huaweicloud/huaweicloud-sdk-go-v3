package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStrategyResponse Response Object
type ListStrategyResponse struct {

	// **参数解释**： 规则实例列表。 **取值范围**： 不涉及。
	Data *[]RuleSet `json:"data,omitempty"`

	// **参数解释**： 规则总数。 **取值范围**： 不涉及。
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStrategyResponse struct{}"
	}

	return strings.Join([]string{"ListStrategyResponse", string(data)}, " ")
}
