package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BandwidthSize 带宽值，单位Mbps。
type BandwidthSize struct {

	// 带宽值，单位Mbps。
	BandwidthSize *int64 `json:"bandwidth_size,omitempty"`
}

func (o BandwidthSize) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthSize struct{}"
	}

	return strings.Join([]string{"BandwidthSize", string(data)}, " ")
}
