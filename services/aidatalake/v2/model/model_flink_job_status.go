package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkJobStatus **参数解释**：Flink作业的状态。
type FlinkJobStatus struct {
	State *FlinkJobState `json:"state,omitempty"`

	// **参数解释**：状态消息详情，包含状态变更的补充说明或错误信息。 **取值范围**：0~1024个字符。
	StateMessage *string `json:"state_message,omitempty"`
}

func (o FlinkJobStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJobStatus struct{}"
	}

	return strings.Join([]string{"FlinkJobStatus", string(data)}, " ")
}
