package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShareTypesAttributionSingleChannel4KLatency struct {

	// 最大值
	Max *int32 `json:"max,omitempty"`

	// 最小值
	Min *int32 `json:"min,omitempty"`
}

func (o ShareTypesAttributionSingleChannel4KLatency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareTypesAttributionSingleChannel4KLatency struct{}"
	}

	return strings.Join([]string{"ShareTypesAttributionSingleChannel4KLatency", string(data)}, " ")
}
