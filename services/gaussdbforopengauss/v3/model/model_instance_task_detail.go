package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceTaskDetail struct {

	// 任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// 任务名称。
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释**: 实例ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o InstanceTaskDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceTaskDetail struct{}"
	}

	return strings.Join([]string{"InstanceTaskDetail", string(data)}, " ")
}
