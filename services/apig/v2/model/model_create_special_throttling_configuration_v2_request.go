package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateSpecialThrottlingConfigurationV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 流控策略的编号

	ThrottleId string `json:"throttle_id"`

	Body *ThrottleSpecialCreate `json:"body,omitempty"`
}

func (o CreateSpecialThrottlingConfigurationV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSpecialThrottlingConfigurationV2Request struct{}"
	}

	return strings.Join([]string{"CreateSpecialThrottlingConfigurationV2Request", string(data)}, " ")
}
