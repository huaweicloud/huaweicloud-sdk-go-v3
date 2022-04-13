package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SSD类型云硬盘预留快照个数，键值对，包含：reserved（预留）、limit（最大）和in_use（已使用）。
type QuotaDetailSnapshotsSsd struct {
	// 已使用的数量。

	InUse int32 `json:"in_use"`
	// 最大的数量。

	Limit int32 `json:"limit"`
	// 预留属性。

	Reserved int32 `json:"reserved"`
}

func (o QuotaDetailSnapshotsSsd) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetailSnapshotsSsd struct{}"
	}

	return strings.Join([]string{"QuotaDetailSnapshotsSsd", string(data)}, " ")
}
