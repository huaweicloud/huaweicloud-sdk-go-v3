package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceBandwidthAutoScalingPolicyResponse Response Object
type DeleteInstanceBandwidthAutoScalingPolicyResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteInstanceBandwidthAutoScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceBandwidthAutoScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceBandwidthAutoScalingPolicyResponse", string(data)}, " ")
}
