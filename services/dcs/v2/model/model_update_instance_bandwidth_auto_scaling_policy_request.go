package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceBandwidthAutoScalingPolicyRequest Request Object
type UpdateInstanceBandwidthAutoScalingPolicyRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *UpdateInstanceBandwidthAutoScalingPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceBandwidthAutoScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceBandwidthAutoScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceBandwidthAutoScalingPolicyRequest", string(data)}, " ")
}
