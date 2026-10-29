package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTasksRulesetsResponse Response Object
type ShowTasksRulesetsResponse struct {

	// **参数解释：** 规则集信息。
	Data *[]TaskRulesetInfo `json:"data,omitempty"`

	// **参数解释：** 总数。 **取值范围:** 不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTasksRulesetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTasksRulesetsResponse struct{}"
	}

	return strings.Join([]string{"ShowTasksRulesetsResponse", string(data)}, " ")
}
