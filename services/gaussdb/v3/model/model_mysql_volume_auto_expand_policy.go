package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlVolumeAutoExpandPolicy struct {

	// **参数解释**：  存储自动扩容上限，需要为10的倍数，单位GB。  **约束限制**：  不涉及。  **取值范围**：  10-128000。  **默认取值**：  不涉及。
	LimitSize *int32 `json:"limit_size,omitempty"`

	// **参数解释**：  可用存储空间率。  **约束限制**：  不涉及。  **取值范围**：  - 1 - 5 - 10 - 15 - 20  **默认取值**：  不涉及。
	TriggerAvailablePercent *int32 `json:"trigger_available_percent,omitempty"`

	// **参数解释**：  扩容步长百分比。  **约束限制**：  不涉及。  **取值范围**：  5-50。  **默认取值**：  不涉及。
	StepPercent *int32 `json:"step_percent,omitempty"`
}

func (o MysqlVolumeAutoExpandPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlVolumeAutoExpandPolicy struct{}"
	}

	return strings.Join([]string{"MysqlVolumeAutoExpandPolicy", string(data)}, " ")
}
