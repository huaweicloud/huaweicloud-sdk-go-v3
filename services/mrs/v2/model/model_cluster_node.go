package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterNode struct {

	// 节点名称，对应manager里的节点名称。
	NodeName *string `json:"node_name,omitempty"`

	// 资源id。确定节点的唯一性，包周期节点可用于计费的查询。
	ResourceId *string `json:"resource_id,omitempty"`

	// 节点组名称。
	NodeGroupName *string `json:"node_group_name,omitempty"`

	// 节点类型。Task、Core、Master等。
	NodeType *string `json:"node_type,omitempty"`

	// on-period包周期或者on-quantity按需。
	BillingType *string `json:"billing_type,omitempty"`

	// 部署类型。支持Server主机类型。
	DeploymentType string `json:"deployment_type"`

	ServerInfo *ServerInfo `json:"server_info,omitempty"`

	CceInfo *CceInfo `json:"cce_info,omitempty"`

	// 节点标签
	Tags *[]Tag `json:"tags,omitempty"`

	NodeDetail *NodeDetail `json:"node_detail,omitempty"`

	// 节点状态。对应页面上的操作状态。
	NodeStatus *string `json:"node_status,omitempty"`

	// 组件实例信息数组。
	ComponentInfos *[]ComponentInfo `json:"component_infos,omitempty"`
}

func (o ClusterNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterNode struct{}"
	}

	return strings.Join([]string{"ClusterNode", string(data)}, " ")
}
