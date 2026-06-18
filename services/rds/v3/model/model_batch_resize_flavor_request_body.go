package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchResizeFlavorRequestBody 批量规格变更请求体请求体
type BatchResizeFlavorRequestBody struct {

	// **参数解释**：   实例ID列表。   **约束限制**：  一次最多下发10个实例。   **取值范围**：  不涉及。  **默认取值**：   不涉及。
	InstanceIds []string `json:"instance_ids"`

	// **参数解释**：  资源规格编码。  **约束限制**：   不涉及。   **取值范围**：  不涉及。   **默认取值**：   不涉及。
	SpecCode string `json:"spec_code"`

	// **参数解释**：  是否进行定时规格变更。  **约束限制**：   不涉及。   **取值范围**：  不涉及。   **默认取值**：   false
	Delay *bool `json:"delay,omitempty"`

	// **参数解释**：  变更包周期实例的规格时可指定，表示是否自动从客户的账户中支付。  **约束限制**：   不涉及。   **取值范围**：  不涉及。   **默认取值**：   false
	AutoPay *bool `json:"auto_pay,omitempty"`

	// **参数解释**：  表示是否占用ip进行规格变更。  **约束限制**：   不涉及。   **取值范围**：  不涉及。   **默认取值**：   true
	OccupyIp *bool `json:"occupy_ip,omitempty"`
}

func (o BatchResizeFlavorRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResizeFlavorRequestBody struct{}"
	}

	return strings.Join([]string{"BatchResizeFlavorRequestBody", string(data)}, " ")
}
