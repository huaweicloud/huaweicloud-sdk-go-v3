package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MysqlAutoScalingRecord 自动变配记录。
type MysqlAutoScalingRecord struct {

	// **参数解释**：  记录ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：  实例ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  变配类型。  **约束限制**：  不涉及。  **取值范围**：  - ENLARGE_FLAVOR：升配 - REDUCE_FLAVOR：降配 - COUNT_UP：只读升配 - COUNT_DOWN：只读降配  **默认取值**：  不涉及。
	ScalingType *string `json:"scaling_type,omitempty"`

	// **参数解释**：  原规格。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	OriginalValue *string `json:"original_value,omitempty"`

	// **参数解释**：  目标规格。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	TargetValue *string `json:"target_value,omitempty"`

	// **参数解释**：  变更结果。  **约束限制**：  不涉及。  **取值范围**：  - SUCCESSFUL：成功 - FAILED：失败  **默认取值**：  不涉及。
	Result *string `json:"result,omitempty"`

	// **参数解释**：  开始时间。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	CreatedAt *int64 `json:"created_at,omitempty"`
}

func (o MysqlAutoScalingRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlAutoScalingRecord struct{}"
	}

	return strings.Join([]string{"MysqlAutoScalingRecord", string(data)}, " ")
}
