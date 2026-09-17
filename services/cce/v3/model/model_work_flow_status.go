package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkFlowStatus struct {
	Phase *WorkFlowPhase `json:"phase,omitempty"`

	// **参数解释：** 升级流程中的各个任务项的执行状态 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PointStatuses *[]PointStatus `json:"pointStatuses,omitempty"`

	// **参数解释：** 表示该升级流程的任务执行线路 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	LineStatuses *[]LineStatus `json:"lineStatuses,omitempty"`
}

func (o WorkFlowStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkFlowStatus struct{}"
	}

	return strings.Join([]string{"WorkFlowStatus", string(data)}, " ")
}
