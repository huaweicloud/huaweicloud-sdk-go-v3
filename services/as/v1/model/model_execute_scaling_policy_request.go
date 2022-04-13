package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExecuteScalingPolicyRequest struct {
	// 伸缩策略ID。

	ScalingPolicyId string `json:"scaling_policy_id"`

	Body *ExecuteScalingPolicyOption `json:"body,omitempty"`
}

func (o ExecuteScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"ExecuteScalingPolicyRequest", string(data)}, " ")
}
