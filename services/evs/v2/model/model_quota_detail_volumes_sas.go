package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SAS类型云硬盘预留的云硬盘个数，键值对，包含：reserved（预留）、limit（最大）和in_use（已使用）。
type QuotaDetailVolumesSas struct {
	// 已使用的数量。

	InUse int32 `json:"in_use"`
	// 最大的数量。

	Limit int32 `json:"limit"`
	// 预留属性。

	Reserved int32 `json:"reserved"`
}

func (o QuotaDetailVolumesSas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetailVolumesSas struct{}"
	}

	return strings.Join([]string{"QuotaDetailVolumesSas", string(data)}, " ")
}
