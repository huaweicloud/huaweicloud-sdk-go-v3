package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceBandwidthAutoScalingPolicyRequest Request Object
type ShowInstanceBandwidthAutoScalingPolicyRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceBandwidthAutoScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceBandwidthAutoScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceBandwidthAutoScalingPolicyRequest", string(data)}, " ")
}
