package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollbackStandardSnapshotRequest Request Object
type RollbackStandardSnapshotRequest struct {

	// **参数解释** 快照ID。 可通过[查询云硬盘快照详情列表](evs_04_2021.xml)接口获取云硬盘快照ID。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	SnapshotId string `json:"snapshot_id"`

	Body *RollbackSnapshotRequestBodyV5 `json:"body,omitempty"`
}

func (o RollbackStandardSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackStandardSnapshotRequest struct{}"
	}

	return strings.Join([]string{"RollbackStandardSnapshotRequest", string(data)}, " ")
}
