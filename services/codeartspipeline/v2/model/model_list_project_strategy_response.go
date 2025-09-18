package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectStrategyResponse Response Object
type ListProjectStrategyResponse struct {

	// **参数解释**： 规则实例列表。 **取值范围**： 不涉及。
	Data *[]RuleSet `json:"data,omitempty"`

	// **参数解释**： 规则总数。 **取值范围**： 不涉及。
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListProjectStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectStrategyResponse struct{}"
	}

	return strings.Join([]string{"ListProjectStrategyResponse", string(data)}, " ")
}
