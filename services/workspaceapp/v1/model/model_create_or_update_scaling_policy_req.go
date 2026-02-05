package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrUpdateScalingPolicyReq 新增或修改弹性伸缩策略。
type CreateOrUpdateScalingPolicyReq struct {

	// 是否启用策略,默认启用： 'true': 启用 'false': 禁用
	Enable *bool `json:"enable,omitempty"`

	// 最大扩容数量。
	MaxScalingAmount int32 `json:"max_scaling_amount"`

	// 单次扩容数量。
	SingleExpansionCount int32 `json:"single_expansion_count"`

	ScalingPolicyBySession *ScalingPolicyBySession `json:"scaling_policy_by_session"`

	ScalingPolicyByResource *ScalingPolicyByResource `json:"scaling_policy_by_resource,omitempty"`

	// 服务器组唯一标识(仅按需服务器组支持该操作)。
	ServerGroupId string `json:"server_group_id"`
}

func (o CreateOrUpdateScalingPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrUpdateScalingPolicyReq struct{}"
	}

	return strings.Join([]string{"CreateOrUpdateScalingPolicyReq", string(data)}, " ")
}
