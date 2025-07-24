package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShareTypesAttributionBandwidth struct {

	// 最大带宽
	Max *int32 `json:"max,omitempty"`

	// 最小带宽
	Min *int32 `json:"min,omitempty"`

	// 带宽梯度
	Step *int32 `json:"step,omitempty"`

	// 带宽密度
	Density float32 `json:"density,omitempty"`

	// 基础带宽
	Base *int32 `json:"base,omitempty"`
}

func (o ShareTypesAttributionBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareTypesAttributionBandwidth struct{}"
	}

	return strings.Join([]string{"ShareTypesAttributionBandwidth", string(data)}, " ")
}
