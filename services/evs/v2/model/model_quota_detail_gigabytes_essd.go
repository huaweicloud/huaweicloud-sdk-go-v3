package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotaDetailGigabytesEssd ESSD类型云硬盘预留的size大小，单位为GiB，键值对，包含：reserved（预留）、limit（最大）和in_use（已使用）。
type QuotaDetailGigabytesEssd struct {

	// 已使用的数量。
	InUse int32 `json:"in_use"`

	// 最大的数量。
	Limit int32 `json:"limit"`

	// 预留属性。
	Reserved int32 `json:"reserved"`
}

func (o QuotaDetailGigabytesEssd) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetailGigabytesEssd struct{}"
	}

	return strings.Join([]string{"QuotaDetailGigabytesEssd", string(data)}, " ")
}
