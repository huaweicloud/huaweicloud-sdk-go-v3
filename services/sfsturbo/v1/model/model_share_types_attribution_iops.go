package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShareTypesAttributionIops struct {

	// 最大iops
	Max *int32 `json:"max,omitempty"`

	// 最小iops
	Min *int32 `json:"min,omitempty"`
}

func (o ShareTypesAttributionIops) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareTypesAttributionIops struct{}"
	}

	return strings.Join([]string{"ShareTypesAttributionIops", string(data)}, " ")
}
