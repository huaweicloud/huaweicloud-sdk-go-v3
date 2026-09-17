package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PointStatus struct {
	TaskType *TaskType `json:"taskType,omitempty"`

	// **参数解释：** 升级任务项ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TaskID *string `json:"taskID,omitempty"`

	Status *UpgradeWorkflowTaskStatus `json:"status,omitempty"`

	// **参数解释：** 升级任务开始时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	StartTimeStamp *string `json:"startTimeStamp,omitempty"`

	// **参数解释：** 升级任务结束时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	EndTimeStamp *string `json:"endTimeStamp,omitempty"`

	// **参数解释：** 升级任务过期时间（当前仅升级前检查任务适用） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ExpireTimeStamp *string `json:"expireTimeStamp,omitempty"`
}

func (o PointStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PointStatus struct{}"
	}

	return strings.Join([]string{"PointStatus", string(data)}, " ")
}
