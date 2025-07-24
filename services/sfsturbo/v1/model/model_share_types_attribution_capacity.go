package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShareTypesAttributionCapacity struct {

	// 最大容量
	Max *int32 `json:"max,omitempty"`

	// 最小容量
	Min *int32 `json:"min,omitempty"`

	// 容量梯度
	Step float32 `json:"step,omitempty"`
}

func (o ShareTypesAttributionCapacity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareTypesAttributionCapacity struct{}"
	}

	return strings.Join([]string{"ShareTypesAttributionCapacity", string(data)}, " ")
}
