package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDetailsOfRequestThrottlingPolicyV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 流控策略的ID

	ThrottleId string `json:"throttle_id"`
}

func (o ShowDetailsOfRequestThrottlingPolicyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfRequestThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfRequestThrottlingPolicyV2Request", string(data)}, " ")
}
