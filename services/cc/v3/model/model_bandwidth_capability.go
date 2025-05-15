package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BandwidthCapability 带宽
type BandwidthCapability struct {

	// 最小带宽。
	Min *int32 `json:"min,omitempty"`

	// 最大带宽。
	Max *int32 `json:"max,omitempty"`
}

func (o BandwidthCapability) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthCapability struct{}"
	}

	return strings.Join([]string{"BandwidthCapability", string(data)}, " ")
}
