package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceNodesInfoResp 实例信息响应体
type InstanceNodesInfoResp struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 当前实例节点总数
	NodeCount *int32 `json:"node_count,omitempty"`

	// 节点详情。
	Nodes *[]NodesInfoResp `json:"nodes,omitempty"`
}

func (o InstanceNodesInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceNodesInfoResp struct{}"
	}

	return strings.Join([]string{"InstanceNodesInfoResp", string(data)}, " ")
}
