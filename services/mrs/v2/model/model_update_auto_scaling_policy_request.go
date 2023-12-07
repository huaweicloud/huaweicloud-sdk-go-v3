package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAutoScalingPolicyRequest Request Object
type UpdateAutoScalingPolicyRequest struct {

	// 集群ID。
	ClusterId string `json:"cluster_id"`

	Body *AutoScalingPolicyV2 `json:"body,omitempty"`
}

func (o UpdateAutoScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAutoScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateAutoScalingPolicyRequest", string(data)}, " ")
}
