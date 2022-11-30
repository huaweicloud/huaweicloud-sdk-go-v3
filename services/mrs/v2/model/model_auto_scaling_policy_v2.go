package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AutoScalingPolicyV2 struct {

	// 节点组名称。
	NodeGroupName *string `json:"node_group_name,omitempty"`

	// 资源计划名称
	ResourcePoolName string `json:"resource_pool_name"`

	AutoScalingPolicy *AutoScalingPolicy `json:"auto_scaling_policy,omitempty"`
}

func (o AutoScalingPolicyV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoScalingPolicyV2 struct{}"
	}

	return strings.Join([]string{"AutoScalingPolicyV2", string(data)}, " ")
}
