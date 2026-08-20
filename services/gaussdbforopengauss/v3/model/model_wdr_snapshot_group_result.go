package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WdrSnapshotGroupResult **参数解释**: 实例内核WDR快照分组详情。
type WdrSnapshotGroupResult struct {

	// **参数解释**: 快照总记录数。 **取值范围**: 不涉及。
	TotalCount int64 `json:"total_count"`

	// **参数解释**: 分组开始时间。对应该分组下的第一个快照的开始时间。 **取值范围**: 不涉及。
	BeginTime int64 `json:"begin_time"`

	// **参数解释**: 分组结束时间。对应该分组下的最后一个快照的结束时间。 **取值范围**: 不涉及。
	EndTime int64 `json:"end_time"`

	// **参数解释**: 分组内快照列表。
	Snapshots []WdrSnapshotRecord `json:"snapshots"`
}

func (o WdrSnapshotGroupResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WdrSnapshotGroupResult struct{}"
	}

	return strings.Join([]string{"WdrSnapshotGroupResult", string(data)}, " ")
}
