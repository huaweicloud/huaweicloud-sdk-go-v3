package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckRulesetParametersResponse Response Object
type CheckRulesetParametersResponse struct {

	// **参数解释：** 历史记录数据。
	Data *[]TaskCheckParameters `json:"data,omitempty"`

	// **参数解释：** 总数。 **取值范围:** ≥0
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CheckRulesetParametersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRulesetParametersResponse struct{}"
	}

	return strings.Join([]string{"CheckRulesetParametersResponse", string(data)}, " ")
}
