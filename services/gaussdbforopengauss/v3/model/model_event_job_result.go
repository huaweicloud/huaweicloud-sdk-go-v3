package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EventJobResult struct {

	// **参数解释**: 事件ID。 **取值范围**: 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**: 实例ID。 **取值范围**: 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**: 任务ID。 **取值范围**: 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**: 错误码。 **取值范围**: 不涉及。
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**: 错误信息。 **取值范围**: 不涉及。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// **参数解释**: 是否成功。 **取值范围**: - true：成功 - false：失败
	Success *bool `json:"success,omitempty"`
}

func (o EventJobResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventJobResult struct{}"
	}

	return strings.Join([]string{"EventJobResult", string(data)}, " ")
}
