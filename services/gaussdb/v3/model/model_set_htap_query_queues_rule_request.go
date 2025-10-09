package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetHtapQueryQueuesRuleRequest Request Object
type SetHtapQueryQueuesRuleRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认值**：  en-us。
	XLanguage string `json:"X-Language"`

	// **参数解释**：  HTAP标准版实例ID，严格匹配UUID规则。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in17，且长度为36个字符。  **默认值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  内容类型。  **约束限制**：  不涉及。  **取值范围**：  application/json。  **默认值**：  application/json。
	ContentType string `json:"Content-Type"`

	Body *OperateHtapQueryQueueRuleReq `json:"body,omitempty"`
}

func (o SetHtapQueryQueuesRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetHtapQueryQueuesRuleRequest struct{}"
	}

	return strings.Join([]string{"SetHtapQueryQueuesRuleRequest", string(data)}, " ")
}
