package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SAS类型云硬盘预留的size大小，单位为GB，键值对，包含：reserved（预留）、limit（最大）和in_use（已使用）。
type QuotaDetailGigabytesSas struct {
	// 已使用的数量。

	InUse int32 `json:"in_use"`
	// 最大的数量。

	Limit int32 `json:"limit"`
	// 预留属性。

	Reserved int32 `json:"reserved"`
}

func (o QuotaDetailGigabytesSas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetailGigabytesSas struct{}"
	}

	return strings.Join([]string{"QuotaDetailGigabytesSas", string(data)}, " ")
}
