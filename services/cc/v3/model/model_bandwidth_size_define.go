package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BandwidthSizeDefine 带宽值，单位Mbps。
type BandwidthSizeDefine struct {
}

func (o BandwidthSizeDefine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthSizeDefine struct{}"
	}

	return strings.Join([]string{"BandwidthSizeDefine", string(data)}, " ")
}
