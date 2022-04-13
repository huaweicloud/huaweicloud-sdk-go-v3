package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 每个云硬盘的容量配额限制。键值对，包含：reserved（预留）、limit（最大）和in_use（已使用）。
type QuotaDetailPerVolumeGigabytes struct {
	// 已使用的数量。

	InUse int32 `json:"in_use"`
	// 最大的数量。

	Limit int32 `json:"limit"`
	// 预留属性。

	Reserved int32 `json:"reserved"`
}

func (o QuotaDetailPerVolumeGigabytes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetailPerVolumeGigabytes struct{}"
	}

	return strings.Join([]string{"QuotaDetailPerVolumeGigabytes", string(data)}, " ")
}
