package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实例信息响应体
type InstanceNodesInfoResp struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// 当前实例节点总数
	NodeCount *int32 `json:"node_count,omitempty" xml:"node_count"`

	// 节点详情。
	Nodes *[]NodesInfoResp `json:"nodes,omitempty" xml:"nodes"`
}

func (o InstanceNodesInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceNodesInfoResp struct{}"
	}

	return strings.Join([]string{"InstanceNodesInfoResp", string(data)}, " ")
}
