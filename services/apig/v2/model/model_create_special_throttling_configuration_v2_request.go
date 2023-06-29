package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSpecialThrottlingConfigurationV2Request Request Object
type CreateSpecialThrottlingConfigurationV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
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
