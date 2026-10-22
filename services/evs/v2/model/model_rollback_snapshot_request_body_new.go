package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollbackSnapshotRequestBodyNew 快照回滚信息。
type RollbackSnapshotRequestBodyNew struct {

	// **参数解释** 回滚的目标云硬盘名称。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释** 回滚的目标云硬盘UUID。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	VolumeId string `json:"volume_id"`
}

func (o RollbackSnapshotRequestBodyNew) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackSnapshotRequestBodyNew struct{}"
	}

	return strings.Join([]string{"RollbackSnapshotRequestBodyNew", string(data)}, " ")
}
