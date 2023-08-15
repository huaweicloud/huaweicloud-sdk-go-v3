package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScalingPolicy 弹性伸缩策略
type ScalingPolicy struct {

	// 是否启用策略,默认启用 'true': 启用 'false': 禁用
	Enable *bool `json:"enable,omitempty"`

	// 最大扩容数量
	MaxScalingAmount int32 `json:"max_scaling_amount"`

	// 单次扩容数量
	SingleExpansionCount int32 `json:"single_expansion_count"`

	ScalingPolicyBySession *ScalingPolicyBySession `json:"scaling_policy_by_session"`
}

func (o ScalingPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingPolicy struct{}"
	}

	return strings.Join([]string{"ScalingPolicy", string(data)}, " ")
}
