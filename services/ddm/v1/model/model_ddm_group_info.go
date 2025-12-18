package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DdmGroupInfo struct {

	// 组ID。
	Id *string `json:"id,omitempty"`

	// 名称。
	Name *string `json:"name,omitempty"`

	// 角色。
	Role *string `json:"role,omitempty"`

	// 组ip。
	Endpoint *string `json:"endpoint,omitempty"`

	// ipv6。
	Ipv6Endpoint *string `json:"ipv6_endpoint,omitempty"`

	// 节点数量。
	NodeCount *int32 `json:"node_count,omitempty"`

	// 负载均衡。
	LoadBalance *bool `json:"load_balance,omitempty"`

	// 是否默认组。
	IsDefaultGroup *bool `json:"is_default_group,omitempty"`

	// 其他信息。
	ExtendMap map[string]string `json:"extend_map,omitempty"`

	// 节点信息。
	Nodes *[]DdmNodeInfo `json:"nodes,omitempty"`
}

func (o DdmGroupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DdmGroupInfo struct{}"
	}

	return strings.Join([]string{"DdmGroupInfo", string(data)}, " ")
}
