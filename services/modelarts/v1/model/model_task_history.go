package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskHistory **参数解释**：训练作业的某次调度的某个实例信息，包含实例IP、实例所在的节点IP、该实例归属于第几次调度等。
type TaskHistory struct {

	// **参数解释**：实例名。 **取值范围**：不涉及。
	Task *string `json:"task,omitempty"`

	// **参数解释**：实例IP。 **取值范围**：不涉及。
	Ip *string `json:"ip,omitempty"`

	// **参数解释**：实例所在的节点IP。 **取值范围**：不涉及。
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**：该实例归属于第几次调度。 **取值范围**：不涉及。
	ScheduleCount *int32 `json:"schedule_count,omitempty"`
}

func (o TaskHistory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskHistory struct{}"
	}

	return strings.Join([]string{"TaskHistory", string(data)}, " ")
}
