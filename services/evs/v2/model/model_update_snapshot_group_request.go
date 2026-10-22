package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSnapshotGroupRequest Request Object
type UpdateSnapshotGroupRequest struct {

	// **参数解释** 快照一致性组ID。 可通过[查询快照一致性组列表](ListSnapshotGroup.xml)接口获取快照一致性组ID。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	SnapshotGroupId string `json:"snapshot_group_id"`

	Body *UpdateSnapshotGroupRequestBody `json:"body,omitempty"`
}

func (o UpdateSnapshotGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSnapshotGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateSnapshotGroupRequest", string(data)}, " ")
}
