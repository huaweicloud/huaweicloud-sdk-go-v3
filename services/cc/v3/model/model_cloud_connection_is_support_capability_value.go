package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloudConnectionIsSupportCapabilityValue 是否支持云连接能力。
type CloudConnectionIsSupportCapabilityValue struct {

	// 是否支持云连接能力。
	IsSupport bool `json:"is_support"`
}

func (o CloudConnectionIsSupportCapabilityValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudConnectionIsSupportCapabilityValue struct{}"
	}

	return strings.Join([]string{"CloudConnectionIsSupportCapabilityValue", string(data)}, " ")
}
