package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EventJobResult struct {

	// **参数解释**：  事件id  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：  实例id  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  任务id  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**：  错误码  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**：  错误信息  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// **参数解释**：  是否成功  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Success *bool `json:"success,omitempty"`
}

func (o EventJobResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventJobResult struct{}"
	}

	return strings.Join([]string{"EventJobResult", string(data)}, " ")
}
