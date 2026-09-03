package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WdrSnapshot WDR快照信息
type WdrSnapshot struct {

	// 快照ID
	SnapshotId *int64 `json:"snapshot_id,omitempty"`

	// 开始时间（Unix timestamp），单位：毫秒
	StartAt *int64 `json:"start_at,omitempty"`

	// 结束时间（Unix timestamp），单位：毫秒
	EndAt *int64 `json:"end_at,omitempty"`
}

func (o WdrSnapshot) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WdrSnapshot struct{}"
	}

	return strings.Join([]string{"WdrSnapshot", string(data)}, " ")
}
