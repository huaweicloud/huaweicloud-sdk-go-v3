package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyHistoryInfo **参数解释**：  参数组应用历史详情  **约束限制**：  不涉及。
type ApplyHistoryInfo struct {

	// **参数解释**：  参数组应用目标实例ID,，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  参数组应用目标实例名称  **约束限制**：  不涉及。  **取值范围**：  不涉及  **默认取值**：  不涉及。
	InstanceName string `json:"instance_name"`

	// **参数解释**：  参数组应用结果。  **约束限制**：  不涉及。  **取值范围**：  不涉及  **默认取值**：  不涉及。
	ApplyResult string `json:"apply_result"`

	// **参数解释**：  参数组应用时间。  **约束限制**：  不涉及。  **取值范围**：  不涉及  **默认取值**：  不涉及。
	ApplyTime string `json:"apply_time"`

	// **参数解释**：  参数组应用错误码  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ErrorCode string `json:"error_code"`
}

func (o ApplyHistoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyHistoryInfo struct{}"
	}

	return strings.Join([]string{"ApplyHistoryInfo", string(data)}, " ")
}
