package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteInstanceBandwidthAutoScalingPolicyResp struct {
}

func (o DeleteInstanceBandwidthAutoScalingPolicyResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceBandwidthAutoScalingPolicyResp struct{}"
	}

	return strings.Join([]string{"DeleteInstanceBandwidthAutoScalingPolicyResp", string(data)}, " ")
}
