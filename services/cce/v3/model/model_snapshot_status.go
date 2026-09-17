package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SnapshotStatus struct {

	// **参数解释：** 任务状态 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Phase *string `json:"phase,omitempty"`

	// **参数解释：** 任务进度 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Progress *string `json:"progress,omitempty"`

	// **参数解释：** 完成时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	CompletionTime *string `json:"completionTime,omitempty"`
}

func (o SnapshotStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotStatus struct{}"
	}

	return strings.Join([]string{"SnapshotStatus", string(data)}, " ")
}
