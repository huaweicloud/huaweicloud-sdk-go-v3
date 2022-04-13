package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 带宽。
type UpdatePublicKibanaBandwidthReqBandWidth struct {
	// 修改后的带宽。

	Size int32 `json:"size"`
}

func (o UpdatePublicKibanaBandwidthReqBandWidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicKibanaBandwidthReqBandWidth struct{}"
	}

	return strings.Join([]string{"UpdatePublicKibanaBandwidthReqBandWidth", string(data)}, " ")
}
