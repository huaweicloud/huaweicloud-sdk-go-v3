package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScalingPolicyResponse Response Object
type CreateScalingPolicyResponse struct {

	// 伸缩策略ID。
	ScalingPolicyId *string `json:"scaling_policy_id,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o CreateScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateScalingPolicyResponse", string(data)}, " ")
}
