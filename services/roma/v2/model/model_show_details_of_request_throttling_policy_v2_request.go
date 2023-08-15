package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDetailsOfRequestThrottlingPolicyV2Request Request Object
type ShowDetailsOfRequestThrottlingPolicyV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 流控策略的编号
	ThrottleId string `json:"throttle_id"`
}

func (o ShowDetailsOfRequestThrottlingPolicyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfRequestThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfRequestThrottlingPolicyV2Request", string(data)}, " ")
}
