package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetHtapQueryQueuesRuleResponse Response Object
type SetHtapQueryQueuesRuleResponse struct {

	// **参数解释**：  响应返回结果。   **取值范围**：  - success：成功。 - failed：失败。
	Status *string `json:"status,omitempty"`

	// **参数解释**：  响应错误信息。  **取值范围**：  不涉及。
	Msg            *string `json:"msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetHtapQueryQueuesRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetHtapQueryQueuesRuleResponse struct{}"
	}

	return strings.Join([]string{"SetHtapQueryQueuesRuleResponse", string(data)}, " ")
}
