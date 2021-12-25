package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateRequestThrottlingPolicyV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 流控策略的编号

	ThrottleId string `json:"throttle_id"`

	Body *ThrottleBaseInfo `json:"body,omitempty"`
}

func (o UpdateRequestThrottlingPolicyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRequestThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"UpdateRequestThrottlingPolicyV2Request", string(data)}, " ")
}
