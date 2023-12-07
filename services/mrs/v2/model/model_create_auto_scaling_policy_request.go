package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAutoScalingPolicyRequest Request Object
type CreateAutoScalingPolicyRequest struct {

	// 集群ID。
	ClusterId string `json:"cluster_id"`

	Body *AutoScalingPolicyV2 `json:"body,omitempty"`
}

func (o CreateAutoScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAutoScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateAutoScalingPolicyRequest", string(data)}, " ")
}
