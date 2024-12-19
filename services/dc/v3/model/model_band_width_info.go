package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BandWidthInfo 带宽信息
type BandWidthInfo struct {

	// 带宽值
	BandwidthSize *int32 `json:"bandwidth_size,omitempty"`

	// 购买全域互连带宽包ID
	GcbId *string `json:"gcb_id,omitempty"`
}

func (o BandWidthInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandWidthInfo struct{}"
	}

	return strings.Join([]string{"BandWidthInfo", string(data)}, " ")
}
