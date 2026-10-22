package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSnapshotGroupRequest Request Object
type DeleteSnapshotGroupRequest struct {

	// **参数解释** 快照一致性组ID。 可通过[查询快照一致性组列表](ListSnapshotGroup.xml)接口获取快照一致性组ID。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	SnapshotGroupId string `json:"snapshot_group_id"`
}

func (o DeleteSnapshotGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSnapshotGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteSnapshotGroupRequest", string(data)}, " ")
}
