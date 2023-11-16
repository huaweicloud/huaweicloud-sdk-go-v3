package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssignedNodeGroup 节点组详情
type AssignedNodeGroup struct {

	// 节点组名称
	Name string `json:"name"`

	// 角色部署信息。 可以指定节点组中部署的角色，该参数是一个字符串数组，每个字符串表示一个角色表达式。 角色表达式定义： - 当该角色在节点组所有节点部署时： {role name}，如“DataNode”。 - 当该角色在节点组指定下标节点部署时：{role name}:{index1},{index2}…,{indexN}，如“NameNode:1,2”，下标从1开始计数。 - 部分角色支持多实例部署（即在一个节点部署多个同角色的实例）：{role name}[{instance count}]，如“EsNode[9]”，多实例部署不需要指定角色位置，默认在节点组所有节点部署多个实例 可选的角色请参考[MRS支持的角色与组件对应表](https://support.huaweicloud.com/api-mrs/mrs_02_0106.html)。
	AssignedRoles []string `json:"assigned_roles"`
}

func (o AssignedNodeGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssignedNodeGroup struct{}"
	}

	return strings.Join([]string{"AssignedNodeGroup", string(data)}, " ")
}
