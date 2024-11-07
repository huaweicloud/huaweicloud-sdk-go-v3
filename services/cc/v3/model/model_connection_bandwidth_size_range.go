package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnectionBandwidthSizeRange json类型
type ConnectionBandwidthSizeRange struct {

	// 最小值
	Min int64 `json:"min"`

	// 最大值
	Max int64 `json:"max"`
}

func (o ConnectionBandwidthSizeRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionBandwidthSizeRange struct{}"
	}

	return strings.Join([]string{"ConnectionBandwidthSizeRange", string(data)}, " ")
}
