package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SpecClusterNode 微服务引擎专享版CCE节点列表
type SpecClusterNode struct {

	// CCE节点信息。
	ClusterNodes *[]ClusterNode `json:"cluster_nodes,omitempty"`
}

func (o SpecClusterNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpecClusterNode struct{}"
	}

	return strings.Join([]string{"SpecClusterNode", string(data)}, " ")
}
