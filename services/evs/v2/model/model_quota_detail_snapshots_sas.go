package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SAS类型云硬盘预留快照个数，键值对，包含：reserved（预留）、limit（最大）和in_use（已使用）。
type QuotaDetailSnapshotsSas struct {

	// 已使用的数量。
	InUse int32 `json:"in_use" xml:"in_use"`

	// 最大的数量。
	Limit int32 `json:"limit" xml:"limit"`

	// 预留属性。
	Reserved int32 `json:"reserved" xml:"reserved"`
}

func (o QuotaDetailSnapshotsSas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetailSnapshotsSas struct{}"
	}

	return strings.Join([]string{"QuotaDetailSnapshotsSas", string(data)}, " ")
}
