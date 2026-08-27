package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventJobResult **参数解释**：  设置事件执行策略响应结果。  **约束限制**：  不涉及。
type EventJobResult struct {

	// **参数解释**：  事件ID。  **取值范围**：  只能由英文字母、数字组成，前面为UUID，后缀为ev07，长度为36个字符。
	EventId *string `json:"event_id,omitempty"`

	// **参数解释**：  实例ID。  **取值范围**：  只能由英文字母、数字组成，前面为UUID，后缀为in07，长度为36个字符。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  创建数据同步的任务ID。  **取值范围**：  不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**：  是否下发成功。  **取值范围**：  - true：下发成功。 - false：下发失败。
	Success *bool `json:"success,omitempty"`

	// **参数解释**：  错误码。  **取值范围**：  不涉及。
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**：  错误消息。  **取值范围**：  不涉及。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o EventJobResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventJobResult struct{}"
	}

	return strings.Join([]string{"EventJobResult", string(data)}, " ")
}
