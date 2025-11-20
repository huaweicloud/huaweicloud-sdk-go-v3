package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobStatus
type JobStatus struct {

	// **参数解释**： 任务的状态 **约束限制**： 不涉及 **取值范围**： - Initializing：初始化 - Running：运行中 - Failed：失败 - Success：成功  **默认取值**： 不涉及
	Phase *string `json:"phase,omitempty"`

	// **参数解释**： 任务变为当前状态的原因 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Reason *string `json:"reason,omitempty"`
}

func (o JobStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobStatus struct{}"
	}

	return strings.Join([]string{"JobStatus", string(data)}, " ")
}
