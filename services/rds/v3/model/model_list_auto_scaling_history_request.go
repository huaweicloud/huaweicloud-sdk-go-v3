package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutoScalingHistoryRequest Request Object
type ListAutoScalingHistoryRequest struct {

	// **参数解释**：  实例ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  查询的变配策略类型。  **约束限制**：  不涉及。  **取值范围**：  - FLAVOR_SCALING：规格变配 - READ_ONLY_SCALING：只读变配  **默认取值**：  FLAVOR_SCALING。
	StrategyType *string `json:"strategy_type,omitempty"`

	// **参数解释**：  索引位置，偏移量。  **约束限制**：  不涉及。  **取值范围**：  不涉及  **默认取值**：  0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：  查询记录数。  **约束限制**：  不涉及。  **取值范围**：  1-100  **默认取值**：  10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAutoScalingHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoScalingHistoryRequest struct{}"
	}

	return strings.Join([]string{"ListAutoScalingHistoryRequest", string(data)}, " ")
}
