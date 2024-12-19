package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceBandwidthAutoScalingPolicyRequest Request Object
type DeleteInstanceBandwidthAutoScalingPolicyRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o DeleteInstanceBandwidthAutoScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceBandwidthAutoScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstanceBandwidthAutoScalingPolicyRequest", string(data)}, " ")
}
