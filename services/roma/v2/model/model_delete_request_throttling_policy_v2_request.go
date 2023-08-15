package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRequestThrottlingPolicyV2Request Request Object
type DeleteRequestThrottlingPolicyV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 流控策略的编号
	ThrottleId string `json:"throttle_id"`
}

func (o DeleteRequestThrottlingPolicyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRequestThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"DeleteRequestThrottlingPolicyV2Request", string(data)}, " ")
}
