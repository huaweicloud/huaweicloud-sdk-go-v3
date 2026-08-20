package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReserveTime 训练作业的保留时长配置。
type ReserveTime struct {

	// **参数解释**：时间单位。  **约束限制**：不涉及。  **取值范围**：  - HOURS：小时   **默认取值**：不涉及。
	TimeUnit string `json:"time_unit"`

	// **参数解释**：保留时长。  **约束限制**：不涉及。  **取值范围**：最小值为1。  **默认取值**：不涉及。
	Duration int32 `json:"duration"`
}

func (o ReserveTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReserveTime struct{}"
	}

	return strings.Join([]string{"ReserveTime", string(data)}, " ")
}
