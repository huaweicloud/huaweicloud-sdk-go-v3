package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkSavepointStatus **参数解释**：Savepoint的状态信息。
type FlinkSavepointStatus struct {
	State *FlinkSavepointState `json:"state,omitempty"`

	// **参数解释**：状态消息详情。 **取值范围**：0~512个字符。
	StateMessage *string `json:"state_message,omitempty"`
}

func (o FlinkSavepointStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkSavepointStatus struct{}"
	}

	return strings.Join([]string{"FlinkSavepointStatus", string(data)}, " ")
}
