package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoScalingPolicyDeleteReq 删除弹性伸缩策略请求。
type AutoScalingPolicyDeleteReq struct {

	// 节点组名称。如果resource_pool_name为default，则删除节点组维度的弹性伸缩策略。如果resource_pool_name不为default，则在该节点组下删除对应资源池维度的策略。
	NodeGroupName string `json:"node_group_name"`

	// 资源池名称。当集群版本不支持按指定资源池进行弹性伸缩时，需要填写为default资源池。不为default时删除指定资源池维度的弹性伸缩策略。
	ResourcePoolName string `json:"resource_pool_name"`
}

func (o AutoScalingPolicyDeleteReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoScalingPolicyDeleteReq struct{}"
	}

	return strings.Join([]string{"AutoScalingPolicyDeleteReq", string(data)}, " ")
}
