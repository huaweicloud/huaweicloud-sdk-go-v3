package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyLabel
type PolicyLabel struct {

	// **参数解释**：转发策略id。  **取值范围**：不涉及
	Id string `json:"id"`

	// **参数解释**：转发策略名称。  **取值范围**：不涉及
	Name string `json:"name"`

	// **参数解释**：转发策略优先级。  **取值范围**：不涉及
	Priority *string `json:"priority,omitempty"`

	// **参数解释**：转发策略action。  **取值范围**：不涉及
	Action string `json:"action"`

	// **参数解释**：转发策略是否启用。  **取值范围**：不涉及
	AdminStateUp bool `json:"admin_state_up"`

	// **参数解释**：规则对象列表。
	Rules []L7Rule `json:"rules"`
}

func (o PolicyLabel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyLabel struct{}"
	}

	return strings.Join([]string{"PolicyLabel", string(data)}, " ")
}
