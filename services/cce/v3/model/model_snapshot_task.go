package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SnapshotTask struct {

	// **参数解释：** 任务类型 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Kind *string `json:"kind,omitempty"`

	// **参数解释：** API版本 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *SnapshotTaskMetadata `json:"metadata,omitempty"`

	Spec *SnapshotSpec `json:"spec,omitempty"`

	Status *SnapshotStatus `json:"status,omitempty"`
}

func (o SnapshotTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotTask struct{}"
	}

	return strings.Join([]string{"SnapshotTask", string(data)}, " ")
}
