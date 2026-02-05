package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrUpdateScalingPolicyResponse Response Object
type CreateOrUpdateScalingPolicyResponse struct {

	// 是否启用策略,默认启用： 'true': 启用 'false': 禁用
	Enable *bool `json:"enable,omitempty"`

	// 最大扩容数量。
	MaxScalingAmount *int32 `json:"max_scaling_amount,omitempty"`

	// 单次扩容数量。
	SingleExpansionCount *int32 `json:"single_expansion_count,omitempty"`

	ScalingPolicyBySession *ScalingPolicyBySession `json:"scaling_policy_by_session,omitempty"`

	ScalingPolicyByResource *ScalingPolicyByResource `json:"scaling_policy_by_resource,omitempty"`
	HttpStatusCode          int                      `json:"-"`
}

func (o CreateOrUpdateScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrUpdateScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateOrUpdateScalingPolicyResponse", string(data)}, " ")
}
