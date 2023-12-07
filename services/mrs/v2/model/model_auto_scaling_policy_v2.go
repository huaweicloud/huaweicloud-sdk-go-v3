package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AutoScalingPolicyV2 struct {

	// 节点组名称。必填参数。如果resource_pool_name为default，则创建节点组维度的弹性伸缩策略。如果resource_pool_name不为default，则在该节点组下创建对应资源池维度的策略。
	NodeGroupName string `json:"node_group_name"`

	// 资源池名称。必填参数。当集群版本不支持按指定资源池进行弹性伸缩时，需要填写为default资源池。不为default时删除指定资源池维度的弹性伸缩策略。
	ResourcePoolName string `json:"resource_pool_name"`

	AutoScalingPolicy *AutoScalingPolicyInfo `json:"auto_scaling_policy,omitempty"`
}

func (o AutoScalingPolicyV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoScalingPolicyV2 struct{}"
	}

	return strings.Join([]string{"AutoScalingPolicyV2", string(data)}, " ")
}
