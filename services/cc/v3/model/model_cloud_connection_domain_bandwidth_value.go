package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloudConnectionDomainBandwidthValue 带宽。
type CloudConnectionDomainBandwidthValue struct {

	// 最小带宽。
	Min *int32 `json:"min,omitempty"`

	// 最大带宽。
	Max *int32 `json:"max,omitempty"`
}

func (o CloudConnectionDomainBandwidthValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudConnectionDomainBandwidthValue struct{}"
	}

	return strings.Join([]string{"CloudConnectionDomainBandwidthValue", string(data)}, " ")
}
