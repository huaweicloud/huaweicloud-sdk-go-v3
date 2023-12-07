package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAutoScalingPolicyRequest Request Object
type DeleteAutoScalingPolicyRequest struct {

	// 集群ID。
	ClusterId string `json:"cluster_id"`

	Body *AutoScalingPolicyDeleteReq `json:"body,omitempty"`
}

func (o DeleteAutoScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAutoScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteAutoScalingPolicyRequest", string(data)}, " ")
}
