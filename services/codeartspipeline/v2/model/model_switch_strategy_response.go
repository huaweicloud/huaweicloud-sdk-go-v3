package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchStrategyResponse Response Object
type SwitchStrategyResponse struct {

	// **参数解释**： 更新状态是否成功。 **取值范围**： - true：状态更新成功。 - false：状态更新失败。
	Status *bool `json:"status,omitempty"`

	// **参数解释**： 更新的策略ID。 **取值范围**： 不涉及。
	RuleSetId      *string `json:"rule_set_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchStrategyResponse struct{}"
	}

	return strings.Join([]string{"SwitchStrategyResponse", string(data)}, " ")
}
