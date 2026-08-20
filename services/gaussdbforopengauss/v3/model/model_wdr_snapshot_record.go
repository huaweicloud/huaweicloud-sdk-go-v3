package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WdrSnapshotRecord **参数解释**: WDR快照记录。
type WdrSnapshotRecord struct {

	// **参数解释**: 快照ID。 **取值范围**: 不涉及。
	Id string `json:"id"`

	// **参数解释**: 快照开始时间。 **取值范围**: 不涉及。
	StartTime int64 `json:"start_time"`

	// **参数解释**: 快照结束时间。 **取值范围**: 不涉及。
	EndTime int64 `json:"end_time"`
}

func (o WdrSnapshotRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WdrSnapshotRecord struct{}"
	}

	return strings.Join([]string{"WdrSnapshotRecord", string(data)}, " ")
}
